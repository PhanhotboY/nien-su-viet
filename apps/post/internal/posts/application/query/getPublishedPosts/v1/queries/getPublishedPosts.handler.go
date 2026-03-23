package queries

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/consumer"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type GetPublishedPostsHandler struct {
	log logger.Logger
}

func NewGetPublishedPostsHandler(
	log logger.Logger,
) consumer.ConsumerHandler {
	return &GetPublishedPostsHandler{
		log: log,
	}
}

func (c *GetPublishedPostsHandler) Handle(
	ctx context.Context, consumeContext types.MessageConsumeContext,
) error {
	c.log.Info("products fetched")

	c.log.Infof("consume context raw data: %v", string(consumeContext.Message().GetData()))
	data := dtoUtil.ValidateConsumeContextData(consumeContext, dto.GetPublicPostsQueryReq{}, c.log)
	c.log.Infof("consume context parsed data: %v", data.Limit)

	return nil
}
