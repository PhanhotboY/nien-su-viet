package rabbitmq_testcontainer

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/rabbitmq/amqp091-go"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type RabbitMQTestContainers struct {
	container testcontainers.Container
	logger    logger.Logger
	cfg       settings.Config
}

func NewRabbitMQTestContainers(cfg settings.Config, l logger.Logger) *RabbitMQTestContainers {
	return &RabbitMQTestContainers{
		cfg:    cfg,
		logger: l,
	}
}

func (r *RabbitMQTestContainers) Start(ctx context.Context, t *testing.T) (settings.RmqConfig, error) {
	r.logger.Info("Starting RabbitMQ test container...")
	containerReq := r.getRunOptions()
	rabbitmqContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return r.cfg.Rmq, err
	}

	r.container = rabbitmqContainer

	host, err := rabbitmqContainer.Host(ctx)
	if err != nil {
		return r.cfg.Rmq, err
	}

	port, err := rabbitmqContainer.MappedPort(ctx, nat.Port(strconv.Itoa(r.cfg.Rmq.Port)))
	if err != nil {
		return r.cfg.Rmq, err
	}

	cfg := settings.RmqConfig{
		QueueName:    r.cfg.Rmq.QueueName,
		DeliveryMode: r.cfg.Rmq.DeliveryMode,
		Persisted:    r.cfg.Rmq.Persisted,
		AppId:        r.cfg.Rmq.AppId,
		AutoStart:    r.cfg.Rmq.AutoStart,
		Reconnecting: r.cfg.Rmq.Reconnecting,
		RmqHostOptions: settings.RmqHostOptions{
			Host:       host,
			Port:       port.Int(),
			UserName:   r.cfg.Rmq.UserName,
			Password:   r.cfg.Rmq.Password,
			RetryDelay: r.cfg.Rmq.RetryDelay,
		},
	}

	isConnectable := IsConnectable(r.logger, cfg)
	if !isConnectable {
		return r.cfg.Rmq, fmt.Errorf("failed to connect to RabbitMQ container at amqp://%s@%s:%d", cfg.UserName, cfg.Host, cfg.Port)
	}

	return cfg, nil
}

func (r *RabbitMQTestContainers) getRunOptions() testcontainers.ContainerRequest {
	containerReq := testcontainers.ContainerRequest{
		Image:        "rabbitmq:4-management-alpine",
		ExposedPorts: []string{strconv.Itoa(r.cfg.Rmq.Port)},
		WaitingFor: wait.ForListeningPort(
			nat.Port(strconv.Itoa(r.cfg.Rmq.Port)),
		).WithStartupTimeout(5 * time.Second),
		HostConfigModifier: func(hostConfig *container.HostConfig) {
			hostConfig.AutoRemove = true
		},
		Hostname: r.cfg.Rmq.Host,
		Env: map[string]string{
			"RABBITMQ_DEFAULT_USER": r.cfg.Rmq.UserName,
			"RABBITMQ_DEFAULT_PASS": r.cfg.Rmq.Password,
		},
	}

	return containerReq
}

func IsConnectable(
	logger logger.Logger,
	options settings.RmqConfig,
) bool {
	conn, err := amqp091.Dial(options.AmqpEndPoint())
	if err != nil {
		logError(
			logger,
			options.UserName,
			options.Password,
			options.Host,
			options.Port,
		)

		return false
	}

	defer conn.Close()

	// https://github.com/michaelklishin/rabbit-hole
	// rmqc, err := rabbithole.NewClient(
	// 	options.HttpEndPoint(),
	// 	options.UserName,
	// 	options.Password,
	// )
	// _, err = rmqc.ListExchanges()

	// if err != nil {
	// 	logger.Errorf(
	// 		"Error in creating rabbitmq connection with http host: %s",
	// 		options.HttpEndPoint(),
	// 	)

	// 	return false
	// }

	logger.Infof(
		"Opened rabbitmq connection on host: amqp://%s@%s:%d",
		options.UserName,
		options.Host,
		options.Port,
	)

	return true
}

func logError(
	logger logger.Logger,
	userName string,
	password string,
	host string,
	hostPort int,
) {
	// we should not use `t.Error` or `t.Errorf` for logging errors because it will `fail` our test at the end and, we just should use logs without error like log.Error (not log.Fatal)
	logger.Errorf(
		"Error in creating rabbitmq connection with amqp host: amqp://%s@%s:%d",
		userName,
		password,
		host,
		hostPort,
	)
}
