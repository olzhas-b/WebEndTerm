package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment   string `envconfig:"ENVIRONMENT" default:"development"`
	IsDevelopment bool   `envconfig:"IS_DEVELOPMENT" default:"true"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`

	DBHost     string `envconfig:"DB_HOST" default:"10.0.85.153"`
	DBPort     int    `envconfig:"DB_PORT" default:"1522"`
	DBName     string `envconfig:"DB_NAME" default:"IBDBTST"`
	DBLogin    string `envconfig:"DB_LOGIN" default:"IBANK"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"For-ibnak12-ibank12"`

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
	m.dsn = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		m.DBHost,
		m.DBPort,
		m.DBLogin,
		m.DBPassword,
		m.DBName,
		"disable",
	)

	return m.dsn
}
