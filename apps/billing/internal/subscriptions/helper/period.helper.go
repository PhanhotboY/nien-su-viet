package subscriptionhelper

import (
	"time"

	planEntity "github.com/phanhotboy/nien-su-viet/apps/billing/internal/plans/domain/entity"
)

func CalculateSubscriptionEndDate(start time.Time, interval planEntity.BillingInterval) time.Time {
	switch interval {
	case planEntity.BILLING_INTERVAL_DAY:
		return start.AddDate(0, 0, 1)
	case planEntity.BILLING_INTERVAL_WEEK:
		return start.AddDate(0, 0, 7)
	case planEntity.BILLING_INTERVAL_MONTH:
		return start.AddDate(0, 1, 0)
	case planEntity.BILLING_INTERVAL_YEAR:
		return start.AddDate(1, 0, 0)
	default:
		return start
	}
}
