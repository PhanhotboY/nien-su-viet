package config

import (
	"reflect"

	sharedCfg "github.com/phanhotboy/nien-su-viet/libs/pkg/config"
)

type PostConfig interface {
	sharedCfg.Config
}

// Fields must be exported for viper mapping
type Config struct {
	sharedCfg.DefaultConfig `mapstructure:",squash"`
}

func ConfigType() reflect.Type {
	return reflect.TypeFor[Config]()
}
