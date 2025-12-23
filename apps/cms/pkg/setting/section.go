package setting

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Postgresql PostgresqlConfig `mapstructure:"postgresql"`
	Rmq        RmqConfig        `mapstructure:"rmq"`
	Security   Security         `mapstructure:"security"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RmqConfig struct {
	Dsn       string `mapstructure:"dsn"`
	QueueName string `mapstructure:"queue_name"`
}

type Security struct {
	AuthCookiePrefix string `mapstructure:"auth_cookie_prefix"`
	BetterAuthSecret string `mapstructure:"better_auth_secret"`
}
