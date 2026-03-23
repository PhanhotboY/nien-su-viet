package settings

type LoggerConfig struct {
	LogLevel      string `mapstructure:"level"`
	CallerEnabled bool   `mapstructure:"callerEnabled"`
	EnableTracing bool   `mapstructure:"enableTracing" default:"true"`
}
