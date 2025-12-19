package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/phanhotboy/nien-su-viet/apps/cms/global"
	"github.com/phanhotboy/nien-su-viet/apps/cms/internal/initialize"
)

func main() {
	r, port := initialize.Run()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started on port", port)

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := global.Amqp.Connection.Close(); err != nil {
		log.Println("Failed to close RabbitMQ connection: ", err)
	}
	if err := global.Amqp.Channel.Close(); err != nil {
		log.Println("Failed to close RabbitMQ channel: ", err)
	}
	if db, err := (global.PostgresDB.DB()); err != nil {
		log.Println("Failed to close Postgres connection: ", err)
	} else {
		if err := db.Close(); err != nil {
			log.Println("Failed to close Postgres DB: ", err)
		}
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown: ", err)
		srv.Close()
	}

	log.Println("Server exiting")
}
