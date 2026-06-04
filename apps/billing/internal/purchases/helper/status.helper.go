package purhelper

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/purchases/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

func ToEntityStatus(status *int32) entity.PurchaseStatus {
	if status == nil {
		return entity.PURCHASE_STATUS_PENDING
	}
	billingStatus := billing_service.PurchaseStatus(*status)
	switch billingStatus {
	case billing_service.PurchaseStatus_PURCHASE_STATUS_PENDING:
		return entity.PURCHASE_STATUS_PENDING
	case billing_service.PurchaseStatus_PURCHASE_STATUS_COMPLETED:
		return entity.PURCHASE_STATUS_COMPLETED
	case billing_service.PurchaseStatus_PURCHASE_STATUS_FAILED:
		return entity.PURCHASE_STATUS_FAILED
	case billing_service.PurchaseStatus_PURCHASE_STATUS_CANCELED:
		return entity.PURCHASE_STATUS_CANCELED
	default:
		return entity.PURCHASE_STATUS_PENDING
	}
}

func ToGrpcStatus(status entity.PurchaseStatus) billing_service.PurchaseStatus {
	switch status {
	case entity.PURCHASE_STATUS_PENDING:
		return billing_service.PurchaseStatus_PURCHASE_STATUS_PENDING
	case entity.PURCHASE_STATUS_COMPLETED:
		return billing_service.PurchaseStatus_PURCHASE_STATUS_COMPLETED
	case entity.PURCHASE_STATUS_FAILED:
		return billing_service.PurchaseStatus_PURCHASE_STATUS_FAILED
	case entity.PURCHASE_STATUS_CANCELED:
		return billing_service.PurchaseStatus_PURCHASE_STATUS_CANCELED
	default:
		return billing_service.PurchaseStatus_PURCHASE_STATUS_PENDING
	}
}
