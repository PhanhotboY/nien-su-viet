package settings

import (
	"strings"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
)

type ServerConfig struct {
	ServiceName string                  `mapstructure:"service_name"`
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
