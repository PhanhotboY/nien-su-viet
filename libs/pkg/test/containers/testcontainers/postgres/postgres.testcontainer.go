package postgres_testcontainer

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	coptions "github.com/phanhotboy/nien-su-viet/libs/pkg/config/options"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresTestContainers struct {
	container testcontainers.Container
	logger    logger.Logger
	cfg       coptions.PostgresqlOptions
}

func NewPostgresTestContainers(c config.Config, l logger.Logger) *PostgresTestContainers {
	return &PostgresTestContainers{
		cfg:    c.GetPostgresqlOptions(),
		logger: l,
	}
}

// func (g *PostgresTestContainers) Cleanup(ctx context.Context) error {
// 	if err := g.container.Terminate(ctx); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (g *PostgresTestContainers) Start(ctx context.Context, t *testing.T) (coptions.PostgresqlOptions, error) {
	g.logger.Info("Starting Postgres test container...")
	containerReq := g.getRunOptions()
	dbContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return g.cfg, err
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		return g.cfg, err
	}
	port, err := dbContainer.MappedPort(ctx, nat.Port(strconv.Itoa(g.cfg.Port)))
	if err != nil {
		return g.cfg, err
	}

	cfg := coptions.PostgresqlOptions{
		Host:     host,
		Port:     port.Int(),
		Username: g.cfg.Username,
		Password: g.cfg.Password,
		Database: g.cfg.Database,
	}

	isConnectable := isConnectable(ctx, g.logger, cfg)
	if !isConnectable {
		return g.cfg, fmt.Errorf("postgres container is not connectable on host: %s:%d", host, cfg.Port)
	}

	g.container = dbContainer

	return cfg, nil
}

func (g *PostgresTestContainers) getRunOptions() testcontainers.ContainerRequest {
	strategies := []wait.Strategy{wait.ForLog("database system is ready to accept connections").
		WithOccurrence(2).
		WithStartupTimeout(5 * time.Second)}
	deadline := 120 * time.Second

	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:alpine",
		ExposedPorts: []string{strconv.Itoa(g.cfg.Port)},
		WaitingFor:   wait.ForAll(strategies...).WithDeadline(deadline),
		Cmd:          []string{"postgres", "-c", "fsync=off"},
		HostConfigModifier: func(hostConfig *container.HostConfig) {
			hostConfig.AutoRemove = true
		},
		Env: map[string]string{
			"POSTGRES_DB":       g.cfg.Database,
			"POSTGRES_PASSWORD": g.cfg.Password,
			"POSTGRES_USER":     g.cfg.Username,
		},
	}

	return containerReq
}

func isConnectable(
	ctx context.Context,
	logger logger.Logger,
	postgresOptions coptions.PostgresqlOptions,
) bool {
	orm, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"postgres://%s:%s@%s:%d/postgres?sslmode=disable",
				postgresOptions.Username,
				postgresOptions.Password,
				postgresOptions.Host,
				postgresOptions.Port,
			),
		),
		&gorm.Config{
			PrepareStmt:              true,
			SkipDefaultTransaction:   true,
			DisableNestedTransaction: true,
		},
	)
	if err != nil {
		logError(logger, postgresOptions.Host, postgresOptions.Port)

		return false
	}

	db, err := orm.DB()
	if err != nil {
		logError(logger, postgresOptions.Host, postgresOptions.Port)

		return false
	}

	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		logError(logger, postgresOptions.Host, postgresOptions.Port)

		return false
	}

	logger.Infof(
		"Opened postgres connection on host: %s:%d", postgresOptions.Host, postgresOptions.Port)

	return true
}

func logError(logger logger.Logger, host string, hostPort int) {
	// we should not use `t.Error` or `t.Errorf` for logging errors because it will `fail` our test at the end and, we just should use logs without error like log.Error (not log.Fatal)
	logger.Errorf("Error in creating postgres connection with %s:%d", host, hostPort)
}
