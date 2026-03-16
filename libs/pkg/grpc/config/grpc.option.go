package config

import (
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/environment"
	typeMapper "github.com/phanhotboy/nien-su-viet/libs/pkg/reflection/typemapper"

	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[GrpcOptions]())

type GrpcOptions struct {
	Port        string `mapstructure:"port"        env:"GRPC_PORT"`
	Host        string `mapstructure:"host"        env:"GRPC_HOST"`
	Development bool   `mapstructure:"development" env:"GRPC_DEVELOPMENT"`
	Name        string `mapstructure:"name"        env:"GRPC_NAME"`
}

func ProvideConfig(environment environment.Environment) (*GrpcOptions, error) {
	return config.BindConfigKey[*GrpcOptions](optionName, environment)
}
