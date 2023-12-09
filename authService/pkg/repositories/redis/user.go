package redis

import "github.com/go-redis/redis"

type RedisUserRepository struct {
	redis *redis.Client
}

func NewRedisUserRepository(redis *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{redis: redis}
}
