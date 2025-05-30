package main

import (
	"github.com/Extremal37/asterisk_db_migrate/internal/app"
	"github.com/Extremal37/asterisk_db_migrate/internal/config"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
)

func main() {
	cfg := config.New()

	err := cfg.ParseConfig()
	if err != nil {
		logger.Fatalf("Failed to parse config: %s", err)
	}

	a := app.New(cfg)

	if err = a.Build(); err != nil {
		logger.Fatalf("Failed to build app: %s", err)
	}

	if err = a.Run(); err != nil {
		logger.Fatalf("Failed to run app: %s", err)
	}

	logger.Info("App successfully stopped")

}
