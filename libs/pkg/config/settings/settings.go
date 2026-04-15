package settings

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Postgresql PostgresqlConfig `mapstructure:"pg"`
	Rmq        RmqConfig        `mapstructure:"rmq"`
	Grpc       GrpcConfig       `mapstructure:"grpc"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	Redis      RedisOptions     `mapstructure:"redis"`
	Metrics    MetricsOptions   `mapstructure:"metrics"`
}

func LoadConfig() Config {
	env := os.Getenv("POST_SERVER_ENV")
	wd, _ := os.Getwd()
	envPath, _ := searchRootDirectory(wd)
	configName := ".env"
	if env == "test" {
		configName = ".env.test"
	}

	v := viper.New()
	v.AddConfigPath(envPath)
	v.SetConfigName(configName)
	v.SetConfigType("env")
	v.SetEnvPrefix("post")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	BindEnvs(v, Config{}, "")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("config file not found: %s", err)
			log.Println("proceeding with environment variables only")
		} else {
			log.Fatalf("error reading config file: %s", err)
		}
	}

	// Watch for configuration changes and reload
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	v.WatchConfig()

	var cfg Config

	if err := v.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.WeaklyTypedInput = true
		dc.Squash = true
	}); err != nil {
		log.Fatalf("unable to decode config into struct: %s", err)
	} else {
		fmt.Printf("Configuration loaded for %s environment\n", v.Get("server.env"))
	}

	return cfg
}

func GetEnv(c Config) environment.Environment {
	return c.Server.Env
}
