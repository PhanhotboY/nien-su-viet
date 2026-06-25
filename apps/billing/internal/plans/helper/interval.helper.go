package planhelper

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

func ToEntityInterval(interval *int32, defaultInterval entity.BillingInterval) entity.BillingInterval {
	if interval == nil {
		return defaultInterval
	}
	billingInterval := billing_service.BillingInterval(*interval)
	switch billingInterval {
	case billing_service.BillingInterval_BILLING_INTERVAL_DAY:
		return entity.BILLING_INTERVAL_DAY
	case billing_service.BillingInterval_BILLING_INTERVAL_WEEK:
		return entity.BILLING_INTERVAL_WEEK
	case billing_service.BillingInterval_BILLING_INTERVAL_MONTH:
		return entity.BILLING_INTERVAL_MONTH
	case billing_service.BillingInterval_BILLING_INTERVAL_YEAR:
		return entity.BILLING_INTERVAL_YEAR
	default:
		return defaultInterval
	}
}

func ToGrpcInterval(interval entity.BillingInterval) billing_service.BillingInterval {
	switch interval {
	case entity.BILLING_INTERVAL_DAY:
		return billing_service.BillingInterval_BILLING_INTERVAL_DAY
	case entity.BILLING_INTERVAL_WEEK:
		return billing_service.BillingInterval_BILLING_INTERVAL_WEEK
	case entity.BILLING_INTERVAL_MONTH:
		return billing_service.BillingInterval_BILLING_INTERVAL_MONTH
	case entity.BILLING_INTERVAL_YEAR:
		return billing_service.BillingInterval_BILLING_INTERVAL_YEAR
	default:
		return billing_service.BillingInterval_BILLING_INTERVAL_MONTH
	}
}
