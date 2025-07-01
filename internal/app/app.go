package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Extremal37/asterisk_db_migrate/internal/config"
	"github.com/Extremal37/asterisk_db_migrate/internal/logger"
	"github.com/Extremal37/asterisk_db_migrate/internal/storage"
	_ "github.com/go-sql-driver/mysql"
)

func New(cfg *config.Config, log *logger.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

type App struct {
	cfg     *config.Config
	log     *logger.Logger
	storage *storage.Storage
}

func (a *App) Run() error {
	a.log.Info("Initializing DB connection...")
	conn, err := initDB(a.cfg)
	if err != nil {
		return fmt.Errorf("failed to init database: %s", err)
	}

	a.storage = storage.New(conn, a.log)
	defer a.storage.Close()

	a.log.Info("Starting migrations...")
	err = a.storage.Migrate(context.Background())
	if err != nil {
		return fmt.Errorf("failed to migrate DB: %w", err)
	}
	return nil
}

func initDB(cfg *config.Config) (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.SourceDB))
	if err != nil {
		return conn, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return conn, fmt.Errorf("failed to ping database: %w", err)
	}

	return conn, nil
}
