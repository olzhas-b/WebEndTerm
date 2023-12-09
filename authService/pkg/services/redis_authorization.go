package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/olzhas-b/PetService/authService/pkg/models"
	"time"
)

type RedisAuthorizationService struct {
	redis       *redis.Client
	tokenConfig models.TokenConfig
}

func NewRedisAuthorizationService(redis *redis.Client, tokenConfig models.TokenConfig) *RedisAuthorizationService {
	return &RedisAuthorizationService{redis: redis, tokenConfig: tokenConfig}
}

func (s *RedisAuthorizationService) Store(key string, token string, t time.Duration) (err error) {
	err = s.redis.Set(key, token, t).Err()
	if err != nil {
		err = fmt.Errorf("RedisAuthorizationService.Store got error: %w", err)
	}
	return err
}

func (s *RedisAuthorizationService) Delete(key string) (err error) {
	err = s.redis.Del(key).Err()
	if err != nil {
		err = fmt.Errorf("RedisAuthorizationService.Delete got error: %w", err)
	}
	return err
}

func (s *RedisAuthorizationService) Get(key string) (token string, err error) {
	val, err := s.redis.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("RedisAuthorizationService.Get got error: %w", err)
	}

	return val, nil
}

func (s *RedisAuthorizationService) Update(key string, value interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}
