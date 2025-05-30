package main

import (
	"github.com/Extremal37/asterisk_db_migrate/internal/config"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
)

func main() {
	cfg := config.New()

	err := cfg.ParseConfig()
	if err != nil {
		logger.Fatalf("failed to build app: %s", err)
	}

	a := app.New(cfg)

	if err = a.Build(); err != nil {
		logger.Fatalf("failed to build app: %s", err)
	}

	if err = a.Run(); err != nil {
		logger.Fatalf("failed to run app: %s", err)
	}

	logger.Info("app successfully stopped")

}
