package queries

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getEvent/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

type getEventHandler struct {
	log          logger.Logger
	tracer       trace.Tracer
	heventEsRepo repository.HistoricalEventEsRepository
	cacheRepo    repository.HistoricalEventCacheRepository
}

type IGetEventHandler interface {
	grpcTypes.GrpcHandler[*GetEventQuery, *dto.GetEventRes]
}

func NewGetEventHandler(
	log logger.Logger,
	heventEsRepo repository.HistoricalEventEsRepository,
	cacheRepo repository.HistoricalEventCacheRepository,
	tracer trace.Tracer,
) IGetEventHandler {
	return getEventHandler{
		log:          log,
		tracer:       tracer,
		heventEsRepo: heventEsRepo,
		cacheRepo:    cacheRepo,
	}
}

func (c getEventHandler) Handle(
	ctx context.Context,
	query *GetEventQuery,
) (*dto.GetEventRes, error) {
	var event *entity.HistoricalEvent
	var err error

	// Try to get event from cache
	event, err = c.cacheRepo.GetHistoricalEvent(ctx, query.ID)
	if err != nil {
		c.log.Warnf("failed to get event from cache: %v, fallback to db", err)
	}

	// If not in cache, fetch from database
	if event == nil {
		event, err = c.heventEsRepo.GetHistoricalEventById(ctx, query.ID)

		if err != nil {
			return nil, grpcerrors.ParseError(err)
		}

		if event == nil {
			return nil, grpcerrors.NewNotFoundErrorGrpcError("event not found", "events.GetEvent")
		}

		// Cache the result
		if err := c.cacheRepo.PutHistoricalEvent(ctx, query.ID, event); err != nil {
			c.log.Warnf("failed to cache event: %v", err)
		}
	}

	res := &dto.GetEventRes{}
	res.FromEntity(event)

	return res, nil
}
