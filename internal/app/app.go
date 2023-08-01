package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"handbooks_backend/config"
	v1 "handbooks_backend/internal/controller/http/v1"
	"handbooks_backend/internal/policy"
	"handbooks_backend/internal/repository"
	"handbooks_backend/internal/service"
	"handbooks_backend/pkg/common/core/identity"
	"handbooks_backend/pkg/common/logging"
	"handbooks_backend/pkg/httpserver"
	"handbooks_backend/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg    *config.Config
	router *gin.Engine
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {
	logging.L(ctx).Info("PostgresSQL initializing")
	// PostgreSQL
	pgConfig := postgres.NewPgConfig(
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)
	pgClient, err := postgres.NewClient(ctx, pgConfig)
	if err != nil {
		logging.L(ctx).Fatal(err)
	}

	logging.L(ctx).Info("Redis initializing")
	// REDIS
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// UUID
	generator := identity.NewGenerator()

	repo := repository.NewRepository(pgClient)
	serv := service.NewService(repo, rdb)
	pol := policy.NewPolicy(serv, generator)

	logging.L(ctx).Info("router initializing")
	router := v1.NewHandler(pol)

	return App{
		cfg:    cfg,
		router: router.Init(),
	}, nil
}

func (a *App) Run(ctx context.Context) {
	httpServer := a.StartHTTP(ctx)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logging.L(ctx).Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logging.L(ctx).Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		logging.L(ctx).Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func (a *App) StartHTTP(ctx context.Context) *httpserver.Server {
	return httpserver.New(
		ctx,
		a.router,
		httpserver.Port(a.cfg.HTTP.Port),
		httpserver.ReadTimeout(a.cfg.HTTP.ReadTimeout),
		httpserver.WriteTimeout(a.cfg.HTTP.WriteTimeout),
	)
}
