package oehelper

import (
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/outbox_events/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/billing_service"
)

func ToEntityStatus(status *int32, defaultStatus entity.OutboxEventStatus) entity.OutboxEventStatus {
	if status == nil {
		return defaultStatus
	}

	switch *status {
	case int32(billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_DEAD):
		return entity.OUTBOX_EVENT_STATUS_DEAD
	case int32(billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_FAILED):
		return entity.OUTBOX_EVENT_STATUS_FAILED
	case int32(billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_PUBLISHED):
		return entity.OUTBOX_EVENT_STATUS_PUBLISHED
	case int32(billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_RETRYING):
		return entity.OUTBOX_EVENT_STATUS_RETRYING
	case int32(billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_PENDING):
		return entity.OUTBOX_EVENT_STATUS_PENDING
	default:
		return defaultStatus
	}
}

func ToGrpcStatus(status entity.OutboxEventStatus) billing_service.OutboxEventStatus {
	switch status {
	case entity.OUTBOX_EVENT_STATUS_DEAD:
		return billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_DEAD
	case entity.OUTBOX_EVENT_STATUS_FAILED:
		return billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_FAILED
	case entity.OUTBOX_EVENT_STATUS_PUBLISHED:
		return billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_PUBLISHED
	case entity.OUTBOX_EVENT_STATUS_RETRYING:
		return billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_RETRYING
	default:
		return billing_service.OutboxEventStatus_OUTBOX_EVENT_STATUS_PENDING
	}
}
