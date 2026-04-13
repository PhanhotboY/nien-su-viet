package settings

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	LogLevel int    `mapstructure:"log_level"`
	SSLMode  bool   `mapstructure:"ssl_mode"`
}
