package fxapp

import (
	"reflect"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/fxapp/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/zap"
	"go.uber.org/fx"
)

type applicationBuilder struct {
	provides  []interface{}
	decorates []interface{}
	options   []fx.Option
	logger    logger.Logger
	config    config.Config
}

func NewApplicationBuilder(cfgType reflect.Type) contracts.ApplicationBuilder {
	cfg := config.LoadConfig(cfgType)

	var logger logger.Logger

	logger = zap.NewZapLogger(cfg)

	return &applicationBuilder{logger: logger, config: cfg}
}

func (a *applicationBuilder) ProvideModule(module fx.Option) {
	a.options = append(a.options, module)
}

func (a *applicationBuilder) Provide(constructors ...interface{}) {
	a.provides = append(a.provides, constructors...)
}

func (a *applicationBuilder) Decorate(constructors ...interface{}) {
	a.decorates = append(a.decorates, constructors...)
}

func (a *applicationBuilder) Build() contracts.Application {
	app := NewApplication(a.provides, a.decorates, a.options, a.logger, a.config)

	return app
}

func (a *applicationBuilder) GetProvides() []interface{} {
	return a.provides
}

func (a *applicationBuilder) GetDecorates() []interface{} {
	return a.decorates
}

func (a *applicationBuilder) Options() []fx.Option {
	return a.options
}

func (a *applicationBuilder) Logger() logger.Logger {
	return a.logger
}

func (a *applicationBuilder) Config() config.Config {
	return a.config
}
