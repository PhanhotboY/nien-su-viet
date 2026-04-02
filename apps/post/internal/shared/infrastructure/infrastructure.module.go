package infrastructure

import (
	"github.com/go-playground/validator"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/core"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/grpc"
	postgres "github.com/phanhotboy/nien-su-viet/libs/pkg/postgresql"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/rabbitmq"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/redis"
	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/
var Module = fx.Module(
	"infrastructureModule",
	// Modules
	// provide core dependencies, e.g., logging, tracing, etc.
	core.Module,
	grpc.Module,
	rabbitmq.Module,
	postgres.Module,
	redis.Module,

	// Other provides
	fx.Provide(validator.New),
)
