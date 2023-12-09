package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/olzhas-b/PetService/authService/pkg/models"
	repo "github.com/olzhas-b/PetService/authService/pkg/repositories"
	"time"
)

type AuthorizationService struct {
	tokenConfig models.TokenConfig
	repo        *repo.Repositories
}

func NewAuthorizationService(repo *repo.Repositories, tokenConfig models.TokenConfig) *AuthorizationService {
	return &AuthorizationService{
		repo:        repo,
		tokenConfig: tokenConfig,
	}
}

func (s *AuthorizationService) GenerateToken(tokenClaim models.TokenClaim) (accessToken, refreshToken string, err error) {
	//TODO implement me
	tokenClaim.Iat = time.Now().Unix()
	tokenClaim.Exp = time.Now().Add(time.Hour * 10).Unix()

	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims["id"] = tokenClaim.ID
	accessTokenClaims["iat"] = tokenClaim.Iat
	accessTokenClaims["exp"] = tokenClaim.Exp
	accessTokenClaims["userType"] = tokenClaim.UserType
	accessTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessToken, err = accessTokenWithClaims.SignedString([]byte(s.tokenConfig.AccessSecret))
	if err != nil {
		err = fmt.Errorf("AuthorizationService.GenerateToken got error: %w", err)
		return
	}

	tokenClaim.Iat = time.Now().Unix()
	tokenClaim.Exp = time.Now().Add(time.Hour * 10).Unix()

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["id"] = tokenClaim.ID
	refreshTokenClaims["iat"] = tokenClaim.Iat
	refreshTokenClaims["exp"] = tokenClaim.Exp
	refreshTokenClaims["userType"] = tokenClaim.UserType
	refreshTokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshToken, err = refreshTokenWithClaims.SignedString([]byte(s.tokenConfig.RefreshSecret))
	return
}

func (s *AuthorizationService) ParseToken(token string, isAccess bool) (tokenClaim models.TokenClaim, err error) {

	JwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to extract token's metadata, unexpected signing method: %v", token.Header["alg"])
		}
		if isAccess {
			return []byte(s.tokenConfig.AccessSecret), nil
		}
		return []byte(s.tokenConfig.RefreshSecret), nil
	})
	if err != nil {
		return
	}

	claims, ok := JwtToken.Claims.(jwt.MapClaims)

	if ok && JwtToken.Valid {
		userID, ok := claims["id"].(float64)
		if !ok {
			return tokenClaim, fmt.Errorf("field id not found")
		}
		tokenClaim.ID = int64(userID)

		iat, ok := claims["iat"].(float64)
		if !ok {
			return tokenClaim, fmt.Errorf("field iat not found")
		}
		tokenClaim.Iat = int64(iat)

		userType, ok := claims["userType"].(string)
		if !ok {
			return tokenClaim, fmt.Errorf("field userType not found")
		}
		tokenClaim.UserType = userType

		exp, ok := claims["exp"].(float64)
		if !ok {
			return tokenClaim, fmt.Errorf("field exp not found")
		}
		expTime := time.Unix(int64(exp), 0)
		if time.Now().After(expTime) {
			return tokenClaim, fmt.Errorf("token expired")
		}
		tokenClaim.Exp = int64(exp)
		return
	}
	return tokenClaim, fmt.Errorf("invalid token")
}
