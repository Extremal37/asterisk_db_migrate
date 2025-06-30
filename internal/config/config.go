package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func New() *Config {
	return &Config{}
}

type Config struct {
	DB
	Log
}

type DB struct {
	Username string `env:"DB_USERNAME" env-default:"asteriskuser"`
	Password string `env:"DB_PASSWORD" env-default:"asteriskpassword"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"3306"`
	SourceDB string `env:"DB_SOURCE_DB" env-default:"asterisk_backup"`
	DestDB   string `env:"DB_DESTINATION_DB" env-default:"asterisk"`
}

type Log struct {
	Level string `env:"LOG_LEVEL" env-default:"INFO"`
}

func (c *Config) ParseConfig() error {
	if err := cleanenv.ReadEnv(c); err != nil {
		return fmt.Errorf("failed to read env's :%w", err)
	}
	return nil
}
