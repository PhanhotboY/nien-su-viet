package initialize

import (
	"log"

	"github.com/phanhotboy/nien-su-viet/apps/blog/global"
	"github.com/phanhotboy/nien-su-viet/apps/blog/pkg/setting"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(cfg *setting.Config) *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	if err := ch.ExchangeDeclare(
		"events",
		"topic",
		true,
		false, false, false, nil,
	); err != nil {
		log.Panicf("Failed to declare exchange: %s", err)
	}

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	ch.QueueBind(q.Name, "user.#", "events", false, nil)
	failOnError(err, "Failed to declare a queue")
	global.Amqp.Channel = ch
	global.Amqp.Connection = conn
	global.Amqp.Queue = q

	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
