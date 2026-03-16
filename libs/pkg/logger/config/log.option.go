package config

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger/models"
	typeMapper "github.com/phanhotboy/nien-su-viet/libs/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[LogOptions]())

type LogOptions struct {
	LogLevel      string         `mapstructure:"level"         env:"LOG_LEVEL"`
	LogType       models.LogType `mapstructure:"logType"       env:"LOG_TYPE"`
	CallerEnabled bool           `mapstructure:"callerEnabled" env:"LOG_CALLER_ENABLED"`
	EnableTracing bool           `mapstructure:"enableTracing" env:"LOG_ENABLE_TRACING" default:"true"`
}

func ProvideLogConfig(env environment.Environment) (*LogOptions, error) {
	return config.BindConfigKey[*LogOptions](optionName, env)
}
