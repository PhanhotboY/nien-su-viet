package settings

import (
	"strings"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
)

type ServerConfig struct {
	ServiceName string                  `mapstructure:"service_name"` // POST_SERVER_SERVICE_NAME
	Version     string                  `mapstructure:"version"`      // POST_SERVER_VERSION
	Host        string                  `mapstructure:"host"`         // POST_SERVER_HOST
	Port        string                  `mapstructure:"port"`         // POST_SERVER_PORT
	Env         environment.Environment `mapstructure:"env"`          // POST_SERVER_ENV
}

func (cfg *ServerConfig) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *ServerConfig) GetMicroserviceName() string {
	return cfg.ServiceName
}
