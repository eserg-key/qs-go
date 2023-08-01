package main

import (
	"context"
	"handbooks_backend/config"
	"handbooks_backend/internal/app"
	"handbooks_backend/pkg/common/logging"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.L(ctx).Info("config initializing")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	ctx = logging.ContextWithLogger(ctx, logging.NewLogger(cfg.Level))

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		logging.L(ctx).Fatal(err)
	}

	logging.L(ctx).Info("Running Application")
	a.Run(ctx)
}
