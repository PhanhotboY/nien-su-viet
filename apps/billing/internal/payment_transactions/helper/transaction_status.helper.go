package pthelper

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/payment_transactions/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

func ToEntityTransactionStatus(s billing_service.PaymentTransactionStatus) entity.PaymentTransactionStatus {
	switch s {
	case billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_PENDING:
		return entity.PAYMENT_TRANSACTION_STATUS_PENDING
	case billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_SUCCEEDED:
		return entity.PAYMENT_TRANSACTION_STATUS_SUCCEEDED
	case billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_FAILED:
		return entity.PAYMENT_TRANSACTION_STATUS_FAILED
	case billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_CANCELLED:
		return entity.PAYMENT_TRANSACTION_STATUS_CANCELLED
	default:
		return entity.PAYMENT_TRANSACTION_STATUS_PENDING
	}
}

func ToGrpcTransactionStatus(s entity.PaymentTransactionStatus) billing_service.PaymentTransactionStatus {
	switch s {
	case entity.PAYMENT_TRANSACTION_STATUS_PENDING:
		return billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_PENDING
	case entity.PAYMENT_TRANSACTION_STATUS_SUCCEEDED:
		return billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_SUCCEEDED
	case entity.PAYMENT_TRANSACTION_STATUS_FAILED:
		return billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_FAILED
	case entity.PAYMENT_TRANSACTION_STATUS_CANCELLED:
		return billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_CANCELLED
	default:
		return billing_service.PaymentTransactionStatus_PAYMENT_TRANSACTION_STATUS_PENDING
	}
}
