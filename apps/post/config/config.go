package config

import (
	"strings"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
)

type Config struct {
	AppOptions AppOptions `mapstructure:"appOptions"`
}

func NewConfig(environment environment.Environment) (*Config, error) {
	cfg, err := config.BindConfig[*Config](environment)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

type AppOptions struct {
	DeliveryType string `mapstructure:"deliveryType" env:"APP_DELIVERY_TYPE"`
	ServiceName  string `mapstructure:"serviceName"   env:"APP_SERVICE_NAME"`
}

func (cfg *AppOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *AppOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
