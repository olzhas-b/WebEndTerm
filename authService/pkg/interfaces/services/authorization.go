package services

import "github.com/olzhas-b/PetService/authService/pkg/models"

type IAuthorizationService interface {
	GenerateToken(tokenClaim models.TokenClaim) (accessToken, refreshToken string, err error)
	ParseToken(token string, isAccess bool) (tokenClaim models.TokenClaim, err error)
}
