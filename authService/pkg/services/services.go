package services

import (
	"github.com/go-redis/redis"
	interfaces "github.com/olzhas-b/PetService/authService/pkg/interfaces/services"
	"github.com/olzhas-b/PetService/authService/pkg/models"
	repo "github.com/olzhas-b/PetService/authService/pkg/repositories"
)

type Services struct {
	interfaces.IUserService
	interfaces.ICipherService
	interfaces.IAuthorizationService
	interfaces.IRedisAuthorizationService
}

func NewServices(repo *repo.Repositories, rConn *redis.Client, tokenConfig models.TokenConfig) *Services {
	return &Services{
		IUserService:               NewUserService(repo, rConn, tokenConfig),
		ICipherService:             NewCipherService(repo, tokenConfig),
		IAuthorizationService:      NewAuthorizationService(repo, tokenConfig),
		IRedisAuthorizationService: NewRedisAuthorizationService(rConn, tokenConfig),
	}
}
