package coptions

import (
	"os"
	"strings"

	cenv "github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	chelper "github.com/phanhotboy/nien-su-viet/libs/pkg/config/helper"
)

type ServerOptions struct {
	ServiceName string           `mapstructure:"service_name"` // PREFIX_SERVER_SERVICE_NAME
	Version     string           `mapstructure:"version"`      // PREFIX_SERVER_VERSION
	Host        string           `mapstructure:"host"`         // PREFIX_SERVER_HOST
	Port        string           `mapstructure:"port"`         // PREFIX_SERVER_PORT
	Env         cenv.Environment `mapstructure:"env"`          // PREFIX_SERVER_ENV
}

func NewDefaultServerOptions() ServerOptions {
	prefix := chelper.GetEnvPrefix()
	return ServerOptions{
		ServiceName: os.Getenv(prefix + "SERVER_SERVICE_NAME"),
		Version:     os.Getenv(prefix + "SERVER_VERSION"),
		Host:        os.Getenv(prefix + "SERVER_HOST"),
		Port:        os.Getenv(prefix + "SERVER_PORT"),
		Env:         cenv.Environment(prefix + os.Getenv("SERVER_ENV")),
	}
}

func (cfg *ServerOptions) GetMicroserviceNameUpper() string {
	return strings.ToUpper(cfg.ServiceName)
}

func (cfg *ServerOptions) GetMicroserviceName() string {
	return cfg.ServiceName
}
