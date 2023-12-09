package services

import (
	"github.com/olzhas-b/PetService/authService/pkg/models"
	"github.com/olzhas-b/PetService/authService/pkg/repositories"
)

type CipherService struct {
	tokenConfig models.TokenConfig
	repo        *repositories.Repositories
}

func NewCipherService(repo *repositories.Repositories, tokenConfig models.TokenConfig) *CipherService {
	return &CipherService{
		repo:        repo,
		tokenConfig: tokenConfig,
	}
}

func (c *CipherService) Decrypt(message string) (decMessage string, err error) {
	//TODO implement me
	panic("implement me")

}

func (c *CipherService) Encrypt(message string) (encMessage string, err error) {
	//TODO implement me
	panic("implement me")
}
