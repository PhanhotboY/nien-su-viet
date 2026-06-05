package coptions

import (
	"os"

	chelper "github.com/phanhotboy/nien-su-viet/libs/pkg/config/helper"
)

type LoggerOptions struct {
	LogLevel      string `mapstructure:"level"`
	CallerEnabled bool   `mapstructure:"caller_enabled"`
	EnableTracing bool   `mapstructure:"enable_tracing" default:"true"`
}

func NewDefaultLoggerOptions() LoggerOptions {
	prefix := chelper.GetEnvPrefix()
	return LoggerOptions{
		LogLevel:      os.Getenv(prefix + "LOGGER_LEVEL"),
		CallerEnabled: os.Getenv(prefix+"LOGGER_CALLER_ENABLED") == "true",
		EnableTracing: os.Getenv(prefix+"LOGGER_ENABLE_TRACING") == "true",
	}
}
