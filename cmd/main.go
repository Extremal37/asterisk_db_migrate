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

	if err = a.Run(); err != nil {
		logger.Fatalf("Failed to run app: %s", err)
	}

	logger.Info("Migration complete!")

}
