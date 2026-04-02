package settings

type RedisOptions struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	Database      int    `mapstructure:"database"`
	PoolSize      int    `mapstructure:"pool.size"`
	EnableTracing bool   `mapstructure:"enable.tracing" default:"true"`
}
