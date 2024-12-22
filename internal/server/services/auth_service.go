package services

import (
	"errors"
	"time"

	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/server/models"
	jwttoken "github.com/besean163/gophkeeper/internal/server/utils/jwt_token"
	"github.com/golang-jwt/jwt/v5"
)

const tokenExpireTime = 1 * time.Hour

type UserRepository interface {
	GetUser(login string) (*models.User, error)
	SaveUser(user *models.User) error
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

func (s AuthService) SaveUser(user *models.User) error {
	return nil
}

func (s AuthService) RegisterUser(login, password string) (string, error) {
	exist, err := s.repository.GetUser(login)
	if err != nil {
		return "", err
	}

	if exist != nil {
		return "", errors.New("user already exist")
	}

	encryptPassword, err := encryptPassword(password)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Login:     login,
		Password:  encryptPassword,
		CreatedAt: time.Now(),
	}

	err = s.repository.SaveUser(user)
	if err != nil {
		return "", err
	}

	return s.createUserToken(user)
}

func (s AuthService) LoginUser(login, password string) (string, error) {
	user, err := s.repository.GetUser(login)
	logger.Get().Println("here")
	logger.Get().Println(user)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exist")
	}

	decryptPassword, err := decryptPassword(user.Password)
	if err != nil {
		return "", err
	}

	if decryptPassword != password {
		return "", errors.New("wrong password")
	}

	return s.createUserToken(user)
}

func (s AuthService) createUserToken(user *models.User) (string, error) {
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
