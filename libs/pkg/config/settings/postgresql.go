package settings

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	LogLevel int    `mapstructure:"log.level"`
	SSLMode  bool   `mapstructure:"ssl.mode"`
}
