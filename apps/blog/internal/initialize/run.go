package initialize

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/phanhotboy/nien-su-viet/apps/blog/global"
)

func Run() {
	LoadConfig()

	_, err := InitPostgreSQL()
	if err != nil {
		panic(err)
	}
	InitRabbitMQ(&global.Config)

	msgs, err := global.Amqp.Channel.Consume(global.Amqp.Queue.Name,
		"", true, false, false, false, nil,
	)
	if err != nil {
		log.Panicf("Failed to register consummer: %v", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received event (%s): %s", msg.RoutingKey, msg.Body)
		}
	}()

	log.Println("Service started. Press Ctrl+C to exit")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Service is shutting down...")

	if err := global.Amqp.Channel.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ channel: %v", err)
	}

	if err := global.Amqp.Connection.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %v", err)
	}

	if db, err := global.PostgresDB.DB(); err == nil {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close PostgreSQL connection: %v", err)
		}
	}

	log.Println("Service shut down gracefully")
}
