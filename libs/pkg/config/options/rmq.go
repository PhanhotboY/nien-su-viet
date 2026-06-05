package coptions

import (
	"fmt"
	"time"
)

type RmqOptions struct {
	RmqHostOptions `mapstructure:",squash"`

	ExchangeName string `mapstructure:"exchange_name"`
	DeliveryMode uint8
	Persisted    bool
	AppId        string
	AutoStart    bool `mapstructure:"autostart"           default:"true"`
	Reconnecting bool `mapstructure:"reconnecting"        default:"true"`
}

type RmqHostOptions struct {
	Host       string    `mapstructure:"hostname"`
	Port       int       `mapstructure:"port"`
	HttpPort   int       `mapstructure:"http_port"`
	UserName   string    `mapstructure:"username"`
	Password   string    `mapstructure:"password"`
	RetryDelay time.Time `mapstructure:"retry_delay"`
}

func (h *RmqHostOptions) AmqpEndPoint() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d", h.UserName, h.Password, h.Host, h.Port)
}

func (h *RmqHostOptions) HttpEndPoint() string {
	return fmt.Sprintf("http://%s:%d", h.Host, h.HttpPort)
}
