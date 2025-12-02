package global

import (
	"github.com/phanhotboy/nien-su-viet/apps/blog/pkg/setting"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

var (
	Config     setting.Config
	PostgresDB *gorm.DB
	Amqp       struct {
		Connection *amqp.Connection
		Channel    *amqp.Channel
		Queue      amqp.Queue
	}
)
