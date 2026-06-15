package adto

import (
	"time"

	"github.com/google/uuid"
	"github.com/phanhotboy/nien-su-viet/apps/billing/internal/subscriptions/domain/entity"
)

type CreateSubscriptionReqDto struct {
	UserID string `json:"user_id" validate:"required"`

	PlanID string `json:"plan_id" validate:"required,uuid"`

	Status int32 `json:"status" validate:"required"`

	// Very common query: find the current subscription that covers now
	CurrentPeriodStart time.Time `json:"current_period_start" validate:"required"`
	CurrentPeriodEnd   time.Time `json:"current_period_end" validate:"required,gtfield=CurrentPeriodStart"`

	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`

	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	ExpiredAt  *time.Time `json:"expired_at,omitempty"`
}

func (d CreateSubscriptionReqDto) MapToEntity() *entity.Subscription {
	return &entity.Subscription{
		UserID: d.UserID,

		PlanID: uuid.MustParse(d.PlanID),

		Status: entity.SubscriptionStatus(d.Status),

		CurrentPeriodStart: d.CurrentPeriodStart,
		CurrentPeriodEnd:   d.CurrentPeriodEnd,

		CancelAtPeriodEnd: d.CancelAtPeriodEnd,

		CanceledAt: d.CanceledAt,
		ExpiredAt:  d.ExpiredAt,
	}
}
