package event

import (
	"encoding/json"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/genproto/events"
)

type PostUpdatedEvent struct {
	types.Message
}

func NewPostUpdatedEvent(post entity.Post) (*PostUpdatedEvent, error) {
	eventData := events.PostCreatedEvent{
		Id:        post.Id.String(),
		Title:     post.Title,
		Slug:      post.Slug,
		Content:   post.Content,
		Summary:   post.Summary,
		Thumbnail: post.Thumbnail,
		AuthorId:  post.AuthorId,
		Views:     int32(post.Views),
		Likes:     int32(post.Likes),
		Published: post.Published,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	}

	if post.PublishedAt != nil {
		dateStr := post.PublishedAt.Format(time.RFC3339)
		eventData.PublishedAt = &dateStr
	}

	data, err := json.Marshal(eventData)
	if err != nil {
		return nil, err
	}
	return &PostUpdatedEvent{
		Message: *types.NewMessage(data),
	}, nil
}
