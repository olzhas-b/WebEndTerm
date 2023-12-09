package services

import (
	"fmt"
	validator "github.com/go-playground/validator/v10"
	"github.com/go-redis/redis"
	"github.com/olzhas-b/PetService/authService/consts"
	interfaces "github.com/olzhas-b/PetService/authService/pkg/interfaces/services"
	"github.com/olzhas-b/PetService/authService/pkg/models"
	"github.com/olzhas-b/PetService/authService/pkg/models/filter"
	"github.com/olzhas-b/PetService/authService/pkg/repositories"
	"log"
)

type UserService struct {
	repo        *repositories.Repositories
	auth        interfaces.IAuthorizationService
	authRedis   interfaces.IRedisAuthorizationService
	tokenConfig models.TokenConfig
}

func NewUserService(repo *repositories.Repositories, rConn *redis.Client, tokenConfig models.TokenConfig) *UserService {
	return &UserService{
		repo:      repo,
		auth:      NewAuthorizationService(repo, tokenConfig),
		authRedis: NewRedisAuthorizationService(rConn, tokenConfig),
	}
}

func (s *UserService) ServiceSignIn(userCred models.UserCredential) (accessToken, refreshToken string, err error) {
	user, err := s.ServiceGetUserByUserCred(userCred)
	if user.ID == 0 || err != nil {
		err = fmt.Errorf("UserService.ServiceSignIn.ServiceGetUserByUserCred got error: %w", err)
		return
	}
	tokenClaim := models.TokenClaim{
		ID:       user.ID,
		UserType: user.Type,
	}

	accessToken, refreshToken, err = s.auth.GenerateToken(tokenClaim)
	if err != nil {
		err = fmt.Errorf("UserService.ServiceSingIn got error: %w", err)
		return
	}

	key := fmt.Sprintf("%s_%d", consts.AccessToken, tokenClaim.ID)
	if err = s.authRedis.Store(key, accessToken, s.tokenConfig.AccessTtl); err != nil {
		err = fmt.Errorf("UserService.ServiceSingIn got error: %w", err)
		return
	}

	key = fmt.Sprintf("%s_%d", consts.RefreshToken, tokenClaim.ID)
	if err = s.authRedis.Store(key, refreshToken, s.tokenConfig.RefreshTtl); err != nil {
		err = fmt.Errorf("UserService.ServiceSingIn got error: %w", err)
		return
	}

	return
}

func (s *UserService) ServiceUpdateToken(token string) (accessToken, refreshToken string, err error) {
	claims, err := s.auth.ParseToken(token, true)
	if err != nil {
		return "", "", err
	}
	if err := s.ServiceLogOut(token); err != nil {
		return "", "", err
	}
	return s.auth.GenerateToken(claims)
}

func (s *UserService) ServiceLogOut(token string) (err error) {
	claims, err := s.auth.ParseToken(token, true)
	if err != nil {
		return fmt.Errorf("UserService.ServiceLogOut got error: %w", err)
	}

	key := fmt.Sprintf("%s_%d", consts.AccessToken, claims.ID)
	if err = s.authRedis.Delete(key); err != nil {
		return fmt.Errorf("UserService.ServiceLogOut got error: %w", err)
	}

	key = fmt.Sprintf("%s_%d", consts.RefreshToken, claims.ID)
	if err = s.authRedis.Delete(key); err != nil {
		return fmt.Errorf("UserService.ServiceLogOut got error: %w", err)
	}

	return
}

func (s *UserService) ServiceGetUserByID(userID int64) (user models.User, err error) {
	return s.repo.GetUserByID(userID)
}

func (s *UserService) ServiceGetUserByUserCred(userCred models.UserCredential) (user models.User, err error) {
	return s.repo.GetUserByUserCred(userCred)
}

func (s *UserService) ServiceCreateUser(user models.User) (result models.User, err error) {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return models.User{}, fmt.Errorf("Validate.Struct: %w", err)
	}
	user.FillDefaultValue()
	result, err = s.repo.CreateUser(user, []string{
		"login",
		"phone",
		"password",
		"first_name",
		"last_name",
	})
	return
}

func (s *UserService) UpdateUser(user models.User) (result models.User, err error) {
	return s.repo.IUserRepository.UpdateUser(user, []string{"todo"})
}

func (s *UserService) ServiceGetAllUsers(f *filter.User) (result []models.User, err error) {
	result, err = s.repo.GetAllUsers(f)
	if err != nil {
		return result, fmt.Errorf("ServiceGetAllUsers : %w", err)
	}
	return
}

func (s *UserService) ServiceCheckToken(token string) (models.TokenClaim, error) {
	tokenClaim, err := s.auth.ParseToken(token, true)
	if err != nil {
		log.Println(err)
		return models.TokenClaim{}, fmt.Errorf("auth.ParseToken got err: %v", err)
	}
	log.Println(tokenClaim)
	key := fmt.Sprintf("%s_%d", consts.AccessToken, tokenClaim.ID)
	if _, err := s.authRedis.Get(key); err != nil {
		log.Println(err)
		return models.TokenClaim{}, fmt.Errorf("token had been expired")
	}

	return tokenClaim, nil
}
