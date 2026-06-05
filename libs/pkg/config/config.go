package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/joho/godotenv"
	cenv "github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	chelper "github.com/phanhotboy/nien-su-viet/libs/pkg/config/helper"
	coptions "github.com/phanhotboy/nien-su-viet/libs/pkg/config/options"
	"github.com/spf13/viper"
)

type Config interface {
	GetEnv() cenv.Environment
	GetServerOptions() coptions.ServerOptions
	GetPostgresqlOptions() coptions.PostgresqlOptions
	GetRmqOptions() coptions.RmqOptions
	GetGrpcOptions() coptions.GrpcOptions
	GetLoggerOptions() coptions.LoggerOptions
	GetRedisOptions() coptions.RedisOptions
	GetMetricsOptions() coptions.MetricsOptions
	GetTracingOptions() coptions.TracingOptions
}

func LoadConfig(cfgType reflect.Type) Config {
	wd, _ := os.Getwd()
	envPath, _ := chelper.SearchRootDirectory(wd)
	fmt.Printf("Loading config from: %s\n", envPath)

	envPrefix := chelper.GetEnvPrefix()
	env := os.Getenv(fmt.Sprintf("%s_SERVER_ENV", envPrefix))
	configName := ".env"
	if env == "test" {
		configName = ".env.test"
	}
	// Load .env file
	err := godotenv.Load(fmt.Sprintf("%s/%s", envPath, configName))
	if err != nil {
		log.Fatalf("Error loading %s file", configName)
	}

	v := viper.New()
	v.AddConfigPath(envPath)
	v.SetConfigName(configName)
	v.SetConfigType("env")
	v.SetEnvPrefix(envPrefix)

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	chelper.BindEnvs(v, reflect.New(cfgType).Elem().Interface(), "")

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

	var res = reflect.New(cfgType).Interface()
	if err := v.Unmarshal(res, func(dc *mapstructure.DecoderConfig) {
		dc.WeaklyTypedInput = true
		dc.Squash = true
	}); err != nil {
		log.Fatalf("unable to decode config into struct: %s", err)
	} else {
		fmt.Printf("Configuration loaded for %s environment\n", v.Get("server.env"))
	}

	return res.(Config)
}
