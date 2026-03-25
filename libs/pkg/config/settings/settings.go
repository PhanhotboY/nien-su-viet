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
	"github.com/subosito/gotenv"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Postgresql PostgresqlConfig `mapstructure:"pg"`
	Rmq        RmqConfig        `mapstructure:"rmq"`
	Grpc       GrpcConfig       `mapstructure:"grpc"`
	Logger     LoggerConfig     `mapstructure:"logger"`
}

type ServerConfig struct {
	ServiceName string                  `mapstructure:"service.name"`
	Host        string                  `mapstructure:"host"`
	Port        string                  `mapstructure:"port"`
	Env         environment.Environment `mapstructure:"env"`
}

func (cfg *ServerConfig) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *ServerConfig) GetMicroserviceName() string {
	return cfg.ServiceName
}

func LoadConfig() Config {
	_ = gotenv.Load(".env")

	viper.SetEnvPrefix("post")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("config file not found: %s", err)
			log.Println("proceeding with environment variables only")
		} else {
			log.Fatalf("error reading config file: %s", err)
		}
	}

	// Watch for configuration changes and reload
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	for _, key := range os.Environ() {
		key = strings.SplitN(key, "=", 2)[0]
		if key, ok := strings.CutPrefix(key, "POST_"); ok {
			key = strings.ToLower(strings.ReplaceAll(key, "_", "."))
			viper.BindEnv(key)
		}
	}

	cfg := Config{}

	if err := viper.Unmarshal(&cfg, func(dc *mapstructure.DecoderConfig) {
		dc.WeaklyTypedInput = true
	}); err != nil {
		log.Fatalf("unable to decode config into struct: %s", err)
	} else {
		fmt.Printf("Configuration loaded for %s environment\n", viper.Get("server.env"))
	}

	return cfg
}

func GetEnv(c Config) environment.Environment {
	return c.Server.Env
}
