package repositories

import (
	"github.com/olzhas-b/PetService/authService/pkg/interfaces/repositories"
	"github.com/olzhas-b/PetService/authService/pkg/repositories/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	repositories.IUserRepository
}

func NewRepositories(DB *gorm.DB) *Repositories {
	return &Repositories{
		IUserRepository: postgres.NewUserRepository(DB),
	}
}
