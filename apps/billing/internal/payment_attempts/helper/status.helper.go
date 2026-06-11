package pahelper

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_attempts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

func ToEntityStatus(status *int32) entity.PaymentAttemptStatus {
	if status == nil {
		return entity.PAYMENT_ATTEMPT_STATUS_CREATED
	}
	billingStatus := billing_service.PaymentAttemptStatus(*status)
	switch billingStatus {
	case billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_CREATED:
		return entity.PAYMENT_ATTEMPT_STATUS_CREATED
	case billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_PENDING:
		return entity.PAYMENT_ATTEMPT_STATUS_PENDING
	case billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_SUCCEEDED:
		return entity.PAYMENT_ATTEMPT_STATUS_SUCCEEDED
	case billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_FAILED:
		return entity.PAYMENT_ATTEMPT_STATUS_FAILED
	case billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_EXPIRED:
		return entity.PAYMENT_ATTEMPT_STATUS_EXPIRED
	case billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_CANCELED:
		return entity.PAYMENT_ATTEMPT_STATUS_CANCELED
	default:
		return entity.PAYMENT_ATTEMPT_STATUS_CREATED
	}
}

func ToGrpcStatus(status entity.PaymentAttemptStatus) billing_service.PaymentAttemptStatus {
	switch status {
	case entity.PAYMENT_ATTEMPT_STATUS_CREATED:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_CREATED
	case entity.PAYMENT_ATTEMPT_STATUS_PENDING:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_PENDING
	case entity.PAYMENT_ATTEMPT_STATUS_SUCCEEDED:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_SUCCEEDED
	case entity.PAYMENT_ATTEMPT_STATUS_FAILED:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_FAILED
	case entity.PAYMENT_ATTEMPT_STATUS_EXPIRED:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_EXPIRED
	case entity.PAYMENT_ATTEMPT_STATUS_CANCELED:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_CANCELED
	default:
		return billing_service.PaymentAttemptStatus_PAYMENT_ATTEMPT_STATUS_CREATED
	}
}
