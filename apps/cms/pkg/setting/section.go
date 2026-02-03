package setting

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Postgresql PostgresqlConfig `mapstructure:"db"`
	Rmq        RmqConfig        `mapstructure:"rmq"`
	Security   Security         `mapstructure:"auth"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RmqConfig struct {
	Dsn       string `mapstructure:"dsn"`
	QueueName string `mapstructure:"queue.name"`
}

type Security struct {
	AuthCookiePrefix string `mapstructure:"cookie.prefix"`
	BetterAuthSecret string `mapstructure:"secret"`
}
