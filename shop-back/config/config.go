package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	Environment   string `envconfig:"ENVIRONMENT" default:"development"`
	IsDevelopment bool   `envconfig:"IS_DEVELOPMENT" default:"true"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`

	DBHost     string `envconfig:"POSTGRES_HOST" default:"0.0.0.0"`
	DBPort     int    `envconfig:"POSTGRES_PORT" default:"5432"`
	DBName     string `envconfig:"POSTGRES_DB" default:"postgres"`
	DBLogin    string `envconfig:"POSTGRES_USER" default:"postgres"`
	DBPassword string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`

	dsn string `envconfig:"DSN" json:"-"`
}

func Load() *Config {
	cfg := new(Config)

	err := envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func (m *Config) DSN() string {
	m.dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))

	return m.dsn
}
