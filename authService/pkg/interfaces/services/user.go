package services

import (
	"github.com/olzhas-b/PetService/authService/pkg/models"
	"github.com/olzhas-b/PetService/authService/pkg/models/filter"
)

type IUserService interface {
	ServiceSignIn(userCred models.UserCredential) (accessToken, refreshToken string, err error)
	ServiceCreateUser(user models.User) (models.User, error)
	ServiceUpdateToken(token string) (accessToken, refreshToken string, err error)
	ServiceLogOut(token string) (err error)
	ServiceGetUserByID(ID int64) (user models.User, err error)
	ServiceGetAllUsers(f *filter.User) ([]models.User, error)
	ServiceGetUserByUserCred(userCred models.UserCredential) (user models.User, err error)
	ServiceCheckToken(token string) (models.TokenClaim, error)
}
