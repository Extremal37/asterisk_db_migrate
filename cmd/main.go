package main

import (
	"fmt"
	"github.com/Extremal37/asterisk_db_migrate/internal/app"
	"github.com/Extremal37/asterisk_db_migrate/internal/config"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
)

func main() {
	cfg := config.New()

	err := cfg.ParseConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to parse config: %s", err))
	}

	l := logger.New(cfg.Level)
	l.Info("Starting...")

	a := app.New(cfg, l)

	if err = a.Run(); err != nil {
		l.Fatalf("Failed to run app: %s", err)
	}

	l.Info("Migration complete!")

}
