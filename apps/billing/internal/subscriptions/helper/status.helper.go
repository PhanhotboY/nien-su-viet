package subscriptionhelper

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

func ToEntityStatus(status *int32) entity.SubscriptionStatus {
	if status == nil {
		return entity.SUBSCRIPTION_STATUS_PENDING
	}
	subscriptionStatus := billing_service.SubscriptionStatus(*status)
	switch subscriptionStatus {
	case billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_EXPIRED:
		return entity.SUBSCRIPTION_STATUS_EXPIRED
	case billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE:
		return entity.SUBSCRIPTION_STATUS_ACTIVE
	case billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_CANCELED:
		return entity.SUBSCRIPTION_STATUS_CANCELED
	case billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_PAST_DUE:
		return entity.SUBSCRIPTION_STATUS_PAST_DUE
	default:
		return entity.SUBSCRIPTION_STATUS_PENDING
	}
}

func ToGrpcStatus(status entity.SubscriptionStatus) billing_service.SubscriptionStatus {
	switch status {
	case entity.SUBSCRIPTION_STATUS_EXPIRED:
		return billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_EXPIRED
	case entity.SUBSCRIPTION_STATUS_ACTIVE:
		return billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_ACTIVE
	case entity.SUBSCRIPTION_STATUS_CANCELED:
		return billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_CANCELED
	case entity.SUBSCRIPTION_STATUS_PAST_DUE:
		return billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_PAST_DUE
	default:
		return billing_service.SubscriptionStatus_SUBSCRIPTION_STATUS_PENDING
	}
}
