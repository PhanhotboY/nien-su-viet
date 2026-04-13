package settings

type LoggerConfig struct {
	LogLevel      string `mapstructure:"level"`
	CallerEnabled bool   `mapstructure:"caller_enabled"`
	EnableTracing bool   `mapstructure:"enable_tracing" default:"true"`
}
