package services

import (
	"time"

	jwttoken "github.com/besean163/gophkeeper/internal/jwt_token"
	"github.com/besean163/gophkeeper/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

const tokenExpireTime = 1 * time.Hour

type UserRepository interface {
	GetUser(login string) (models.User, error)
}

type AuthService struct {
	secret     string
	repository UserRepository
}

func NewAuthService(secret string, repository UserRepository) AuthService {
	return AuthService{
		secret:     secret,
		repository: repository,
	}
}

func (s AuthService) GetUser(login string) *models.User {
	return &models.User{}
}

func (s AuthService) SaveUser(user *models.User) error {
	return nil
}

func (s AuthService) CreateUser(user *models.User) {

}

func (s AuthService) RegisterUser(user *models.User) (string, error) {
	s.SaveUser(user)
	return s.CreateUserToken(user)
}

func (s AuthService) CreateUserToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Login,
		"exp":     time.Now().Add(tokenExpireTime).Unix(), // Время истечения
	}
	tokenString, err := jwttoken.GetJWTToken(s.secret, claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
