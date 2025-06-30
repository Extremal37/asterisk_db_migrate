package main

import (
	"github.com/Extremal37/asterisk_db_migrate/internal/app"
	"github.com/Extremal37/asterisk_db_migrate/internal/config"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
)

func main() {
	l := logger.New()
	l.Info("Starting...")

	cfg := config.New()

	err := cfg.ParseConfig()
	if err != nil {
		l.Fatalf("Failed to parse config: %s", err)
	}

	a := app.New(cfg, l)

	if err = a.Run(); err != nil {
		l.Fatalf("Failed to run app: %s", err)
	}

	l.Info("Migration complete!")

}
