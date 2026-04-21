package settings

type OTLPProvider struct {
	OTLPEndpoint string `mapstructure:"endpoint"` // POST_METRICS_OTEL_COLLECTOR_ENDPOINT
}

type MetricsOptions struct {
	UseStdout            bool         `mapstructure:"use_stdout"`           // POST_METRICS_USE_STDOUT
	InstrumentationName  string       `mapstructure:"instrumentation_name"` // POST_METRICS_INSTRUMENTATION_NAME
	EnableHostMetrics    bool         `mapstructure:"enable_host_metrics"`  // POST_METRICS_ENABLE_HOST_METRICS
	UseOTLP              bool         `mapstructure:"use_otlp"`             // POST_METRICS_USE_OTLP
	OtelCollectorOptions OTLPProvider `mapstructure:"otlp_provider"`
}
