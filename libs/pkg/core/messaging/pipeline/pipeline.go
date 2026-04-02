package pipeline

import (
	"context"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/core/messaging/types"
)

type ConsumerHandlerFunc func(ctx context.Context) error

type ConsumerPipeline interface {
	Handle(ctx context.Context, consumerContext types.MessageConsumeContext, next ConsumerHandlerFunc) error
}

type ConsumerPipelineConfiguration struct {
	Pipelines []ConsumerPipeline
}

type ConsumerPipelineConfigurationBuilder interface {
	AddPipeline(pipeline ConsumerPipeline) ConsumerPipelineConfigurationBuilder
	Build() *ConsumerPipelineConfiguration
}

type ConsumerPipelineConfigurationBuilderFunc func(builder ConsumerPipelineConfigurationBuilder)

type consumerPipelineConfigurationBuilder struct {
	pipelines []ConsumerPipeline
}

func NewConsumerPipelineConfigurationBuilder() ConsumerPipelineConfigurationBuilder {
	return &consumerPipelineConfigurationBuilder{pipelines: []ConsumerPipeline{}}
}

func (b *consumerPipelineConfigurationBuilder) AddPipeline(pipeline ConsumerPipeline) ConsumerPipelineConfigurationBuilder {
	b.pipelines = append(b.pipelines, pipeline)
	return b
}

func (b *consumerPipelineConfigurationBuilder) Build() *ConsumerPipelineConfiguration {
	return &ConsumerPipelineConfiguration{Pipelines: b.pipelines}
}
