package postgres

import (
	"github.com/olzhas-b/PetService/authService/pkg/models"
	"github.com/olzhas-b/PetService/authService/pkg/models/filter"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (repo *UserRepository) GetUserByID(ID int64) (user models.User, err error) {
	err = repo.DB.Model(models.User{}).
		Where("id = ?", ID).
		First(&user).
		Error
	return
}

func (repo *UserRepository) GetAllUsers(filter *filter.User) (users []models.User, err error) {
	err = repo.DB.Model(models.User{}).
		Find(&users).
		Error
	return
}

func (repo *UserRepository) UpdateUser(user models.User, selectedColumns []string) (result models.User, err error) {
	err = repo.DB.Model(models.User{}).
		Select(selectedColumns).
		Save(user).
		Error
	if err != nil {
		return
	}
	return repo.GetUserByID(user.ID)
}

func (repo *UserRepository) CreateUser(user models.User, selectedColumns []string) (models.User, error) {
	err := repo.DB.Model(models.User{}).
		Select(selectedColumns).
		Create(&user).
		Error

	return user, err
}

func (repo *UserRepository) GetUserByUserCred(userCred models.UserCredential) (result models.User, err error) {
	err = repo.DB.Model(models.User{}).
		Where("login = (?)", userCred.Login).
		Where("password = (?)", userCred.Password).
		First(&result).
		Error
	return
}
