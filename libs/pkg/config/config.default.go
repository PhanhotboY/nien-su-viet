package config

import (
	"reflect"

	cenv "github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	coptions "github.com/phanhotboy/nien-su-viet/libs/pkg/config/options"
)

// Fields must be exported for viper mapping
type DefaultConfig struct {
	Server     coptions.ServerOptions     `mapstructure:"server"`
	Postgresql coptions.PostgresqlOptions `mapstructure:"pg"`
	Rmq        coptions.RmqOptions        `mapstructure:"rmq"`
	Grpc       coptions.GrpcOptions       `mapstructure:"grpc"`
	Logger     coptions.LoggerOptions     `mapstructure:"logger"`
	Redis      coptions.RedisOptions      `mapstructure:"redis"`
	Metrics    coptions.MetricsOptions    `mapstructure:"metrics"`
	Tracing    coptions.TracingOptions    `mapstructure:"tracing"`
}

func NewDefaultConfig() Config {
	return LoadConfig(reflect.TypeFor[DefaultConfig]())
}

func (c DefaultConfig) GetEnv() cenv.Environment {
	return c.Server.Env
}

func (c DefaultConfig) GetServerOptions() coptions.ServerOptions {
	return c.Server
}

func (c DefaultConfig) GetPostgresqlOptions() coptions.PostgresqlOptions {
	return c.Postgresql
}

func (c DefaultConfig) GetRmqOptions() coptions.RmqOptions {
	return c.Rmq
}

func (c DefaultConfig) GetGrpcOptions() coptions.GrpcOptions {
	return c.Grpc
}

func (c DefaultConfig) GetLoggerOptions() coptions.LoggerOptions {
	return c.Logger
}

func (c DefaultConfig) GetRedisOptions() coptions.RedisOptions {
	return c.Redis
}

func (c DefaultConfig) GetMetricsOptions() coptions.MetricsOptions {
	return c.Metrics
}

func (c DefaultConfig) GetTracingOptions() coptions.TracingOptions {
	return c.Tracing
}
