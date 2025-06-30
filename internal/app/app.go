package app

import (
	"database/sql"
	"fmt"
	"github.com/Extremal37/asterisk_db_migrate/internal/config"
	"github.com/Extremal37/asterisk_db_migrate/internal/storage"
	_ "github.com/go-sql-driver/mysql"
)

func New(cfg *config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

type App struct {
	cfg     *config.Config
	storage *storage.Storage
}

func (a *App) Run() error {
	conn, err := initDB(a.cfg)
	if err != nil {
		return fmt.Errorf("failed to init database: %s", err)
	}

	a.storage = storage.New(conn)

	err = a.storage.Migrate()
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
