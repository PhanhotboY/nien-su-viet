package initialize

import (
	"fmt"
	"os"

	"github.com/phanhotboy/nien-su-viet/apps/blog/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.AddConfigPath("./configs")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found: %w", err))
		}
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}
}
