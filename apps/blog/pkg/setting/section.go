package setting

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Grpc       GrpcConfig       `mapstructure:"grpc"`
	Postgresql PostgresqlConfig `mapstructure:"postgresql"`
	Security   Security         `mapstructure:"security"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type GrpcConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type PostgresqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Security struct {
	Jwt struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}
