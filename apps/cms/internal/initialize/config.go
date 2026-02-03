package initialize

import (
	"fmt"
	"log"
	"os"
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
		if key, ok := strings.CutPrefix(key, "CMS_"); ok {
			key = strings.ToLower(strings.ReplaceAll(key, "_", "."))
			viper.BindEnv(key)
		}
	}

	if err := viper.Unmarshal(&global.Config, func(dc *mapstructure.DecoderConfig) {
		dc.WeaklyTypedInput = true
	}); err != nil {
		log.Fatalf("unable to decode config into struct: %s", err)
	} else {
		fmt.Printf("Configuration loaded for %s environment\n", viper.Get("server.env"))
	}
}
