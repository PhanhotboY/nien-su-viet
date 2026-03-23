package global

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

var (
	PostgresDB *gorm.DB
	Amqp       struct {
		Connection *amqp.Connection
		Channel    *amqp.Channel
		Queue      amqp.Queue
	}
)
