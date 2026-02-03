package initialize

import (
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-viper/mapstructure/v2"
	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func LoadConfig() {
	_ = gotenv.Load(".env")

	viper.SetEnvPrefix("cms")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("config file not found: %s", err)
		}
		log.Fatalf("error reading config file: %s", err)
	}

	// Watch for configuration changes and reload
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	// Bind environment variables to nested configuration keys
	for _, key := range viper.AllKeys() {
		viper.BindEnv(strings.Join(strings.Split(key, "_")[1:], "."))
	}

	if err := viper.Unmarshal(&global.Config, func(dc *mapstructure.DecoderConfig) {
		dc.WeaklyTypedInput = true
	}); err != nil {
		log.Fatalf("unable to decode config into struct: %s", err)
	} else {
		fmt.Printf("Configuration loaded for %s environment\n", viper.Get("server.env"))
	}
}
