package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	Debug    bool
	Redis    RedisConfig
	Database DatabaseConfig
	Token    TokenConfig
	HTTP     HTTPConfig
	TimeOut  bool
}

type HTTPConfig struct {
	Http string
	Name string
	Port string
}

func (h *HTTPConfig) DNS() string {
	return fmt.Sprintf("%s://%s:%s", h.Http, h.Name, h.Port)
}

type RedisConfig struct {
	Enable   bool
	Addr     string
	Port     string
	Password string
	DB       int
}

type DatabaseConfig struct {
	Enable   bool
	Driver   string
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Dsn      string
}

type TokenConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTtl     time.Duration
	RefreshTtl    time.Duration
}

var (
	globalConfig = Config{}
)

func InitConfig() {
	globalConfig = Config{
		Token: TokenConfig{
			AccessSecret:  "asdflsadaqjwe123DEavlkjl12312312",
			RefreshSecret: "fadsf0ivoi@vlka0sd123,vk234/adsf;1!1231$$$#123",
			AccessTtl:     time.Hour * 1000000,
			RefreshTtl:    time.Hour * 1000000,
		},
		HTTP: HTTPConfig{
			Name: "0.0.0.0",
			Port: "8080",
			Http: "http",
		},
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS-ADDR"),
			Password: os.Getenv("REDIS-PASSWORD"),
		},
	}
	log.Println(os.Getenv("REDIS-ADDR"))
}

func Get() *Config {
	return &globalConfig
}
