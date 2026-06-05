package coptions

type GrpcOptions struct {
	Port        string `mapstructure:"port"`
	Host        string `mapstructure:"host"`
	Development bool   `mapstructure:"development"`
	Name        string `mapstructure:"name"`
}
