package acmd

import (
	"context"
	"fmt"
	"time"

	adto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/application/commands/createPurchase/dto"
	drepo "github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/repository"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/config"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/shared/infrastructure/zalopay"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dbcontracts "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"

	createPaymentAttempt "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt"
	createPaymentAttemptDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/commands/createPaymentAttempt/dto"

	getPlanById "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById"
	getPlanByIdDto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/application/queries/getPlanById/dto"

	getPAQuery "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt"
	getPADto "github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/application/queries/getPaymentAttempt/dto"
)

type CreatePurchaseHandler interface {
	grpcTypes.GrpcHandler[*CreatePurchaseCommand, adto.CreatePurchaseResDto]
}

type createPurchaseHandler struct {
	logger   logger.Logger
	purCache drepo.PurchaseCacheRepo
	purDb    drepo.PurchaseDBRepo
	zpClient *zalopay.Client
	redis    redis.RedisClientWithExpire
	db       dbcontracts.TxContextDb
	cfg      config.BillingConfig

	createPaymentAttemptHandler createPaymentAttempt.CreatePaymentAttemptHandler
	getPlanByIdHandler          getPlanById.GetPlanByIdHandler
	getPAHandler                getPAQuery.GetPaymentAttemptHandler
}

func NewCreatePurchaseHandler(
	l logger.Logger,
	db dbcontracts.TxContextDb,
	purCache drepo.PurchaseCacheRepo,
	purDb drepo.PurchaseDBRepo,
	zpClient *zalopay.Client,
	r redis.RedisClientWithExpire,
	cfg config.BillingConfig,

	createPaymentAttemptHandler createPaymentAttempt.CreatePaymentAttemptHandler,
	getPlanByIdHandler getPlanById.GetPlanByIdHandler,
	getPAHandler getPAQuery.GetPaymentAttemptHandler,
) CreatePurchaseHandler {
	return &createPurchaseHandler{
		logger:   l,
		purCache: purCache,
		purDb:    purDb,
		zpClient: zpClient,
		redis:    r,
		db:       db,
		cfg:      cfg,

		createPaymentAttemptHandler: createPaymentAttemptHandler,
		getPlanByIdHandler:          getPlanByIdHandler,
		getPAHandler:                getPAHandler,
	}
}

func (h *createPurchaseHandler) Handle(ctx context.Context, command *CreatePurchaseCommand) (adto.CreatePurchaseResDto, error) {
	// Set idempotency key to prevent duplicate processing
	if err := h.redis.SetNX(ctx, command.IdempotencyKey, "processing", time.Minute*5); err != nil {
		h.logger.Errorf("Failed to set idempotency key: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to process purchase: 1", "NewCreatePurchaseHandler")
	}
	// Clean up idempotency key after processing
	defer h.redis.Del(ctx, command.IdempotencyKey)

	existingPurchase, err := h.purDb.GetPurchaseByIdempotencyKey(ctx, command.IdempotencyKey)
	if err != nil {
		h.logger.Errorf("Failed to check existing purchase by idempotency key: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to process purchase: 2", "NewCreatePurchaseHandler")
	}
	if existingPurchase != nil {
		h.logger.Infof("Purchase already exists for idempotency key: %s", command.IdempotencyKey)
		return nil, grpcerrors.NewBadRequestGrpcError("Duplicated Request", "NewCreatePurchaseHandler")
	}

	// TODO: Start transaction
	var purchaseId string
	var paymentRes *zalopay.CreateOrderResponse
	var paymentAttemptRes getPADto.GetPaymentAttemptResDto
	err = h.db.RunInTx(ctx, func(ctx context.Context, tx dbcontracts.TxContextDb) error {
		// TODO: fetch plan details
		getPlanByIdQuery, _ := getPlanById.NewGetPlanByIdQuery(&getPlanByIdDto.GetPlanByIdReqDto{
			Id: command.PlanID,
		})
		planRes, err := h.getPlanByIdHandler.Handle(ctx, getPlanByIdQuery)
		if err != nil {
			h.logger.Errorf("Failed to fetch plan details: %v", err)
			return err
		}

		plan := planRes.GetData()
		purchaseEntity := command.MapToEntity()
		purchaseEntity.Amount = plan.Price.Amount
		purchaseEntity.Currency = plan.Price.Currency
		purchaseId, err = h.purDb.CreatePurchase(ctx, purchaseEntity)
		if err != nil {
			h.logger.Errorf("Failed to create purchase: %v", err)
			return err
		}

		appTransID := zalopay.GenerateAppTransID()
		paymentRes, err = h.zpClient.CreateOrder(ctx, zalopay.CreateOrderRequest{
			AppTime:    time.Now().UnixMilli(),
			AppTransID: appTransID,
			AppUser:    command.UserID,
			// Use ZaloPay payment gateway
			BankCode: "",

			// TODO: fetch plan and update fields
			Amount: plan.Price.Amount,
			Title:  fmt.Sprintf("Thanh toán gói thành viên: %s", plan.Name),
			Item: zalopay.Item{
				ItemID:    plan.Id.String(),
				ItemName:  plan.Name,
				ItemPrice: plan.Price.Amount,
			}.String(),
			Description: fmt.Sprintf("Thanh toan goi thanh vien %s", plan.Name),
			// Temporary no data
			EmbedData: zalopay.EmbedData{
				PurchaseID:  purchaseId,
				PlanID:      plan.Id.String(),
				RedirectURL: h.cfg.GetZaloPayOptions().RedirectURL,
			}.String(),
			CallbackURL: h.cfg.GetZaloPayOptions().CallbackURL,
		})
		if err != nil {
			h.logger.Errorf("Failed to create order with ZaloPay: %v", err)
			return err
		}
		if paymentRes.ReturnCode != 1 {
			h.logger.Errorf("ZaloPay create order failed: %s", paymentRes.SubReturnMessage)
			return fmt.Errorf("ZaloPay create order failed: %s", paymentRes.ReturnMessage)
		}

		// TODO: Create PENDING payment attempt
		createPaymentAttemptCmd, err := createPaymentAttempt.NewCreatePaymentAttemptCommand(
			&createPaymentAttemptDto.CreatePaymentAttemptReqDto{
				PurchaseID:            purchaseId,
				Provider:              "zalopay",
				ProviderTransactionID: appTransID,
				CheckoutURL:           paymentRes.OrderURL,
				Amount:                plan.Price,
				ProviderMetadata:      "{}",
			})
		if err != nil {
			h.logger.Errorf("Failed to create CreatePaymentAttemptCommand: %v", err)
			return err
		}
		createPARes, err := h.createPaymentAttemptHandler.Handle(ctx, createPaymentAttemptCmd)
		if err != nil {
			return err
		}
		getPAByPurchaseIDQuery, err := getPAQuery.NewGetPaymentAttemptQuery(
			&getPADto.GetPaymentAttemptReqDto{
				PaymentAttemptId: createPARes.GetData().ID,
			})
		if err != nil {
			h.logger.Errorf("Failed to create GetPaymentAttemptByPurchaseIDQuery: %v", err)
			return err
		}
		paymentAttemptRes, err = h.getPAHandler.Handle(ctx, getPAByPurchaseIDQuery)
		if err != nil {
			h.logger.Errorf("Failed to get payment attempt after creation: %v", err)
			return err
		}

		// Invalidate cache after creating purchase
		err = h.purCache.DeleteAllPurchases(ctx)
		if err != nil {
			// Todo: Handle cache deletion failure
			h.logger.Warnf("Failed to delete all purchases cache after creating purchase: %v", err)
		}

		return nil
	})
	if err != nil {
		h.logger.Errorf("Transaction failed: %v", err)
		return nil, grpcerrors.NewInternalServerGrpcError("Failed to process purchase: 4 "+err.Error(), "NewCreatePurchaseHandler")
	}

	return adto.NewCreatePurchaseResDto(purchaseId, paymentRes, paymentAttemptRes), nil
}
