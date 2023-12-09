package database

import (
	"github.com/go-redis/redis"
	"github.com/olzhas-b/PetService/authService/config"
	"log"
)

func InitRedis() (*redis.Client, error) {
	conf := config.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: "",
		DB:       2,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return client, err
	}
	log.Println(pong)

	return client, err
}
