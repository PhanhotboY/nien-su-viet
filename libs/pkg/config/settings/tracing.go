package settings

type TracingOptions struct {
	Enabled             bool         `mapstructure:"enabled"`              // POST_TRACING_ENABLED
	InstrumentationName string       `mapstructure:"instrumentation_name"` // POST_TRACING_INSTRUMENTATION_NAME
	AlwaysOnSampler     bool         `mapstructure:"always_on_sampler"`    // POST_TRACING_ALWAYS_ON_SAMPLER
	UseStdout           bool         `mapstructure:"use_stdout"`           // POST_TRACING_USE_STDOUT
	UseOTLP             bool         `mapstructure:"use_otlp"`             // POST_TRACING_USE_OTLP
	OtlpProvider        OTLPProvider `mapstructure:"otlp_provider"`
}
