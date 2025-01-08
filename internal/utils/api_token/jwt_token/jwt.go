package jwttoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/golang-jwt/jwt/v5"
)

const tokenExpireTime = 1 * time.Hour

type JWTTokener struct {
	secret string
}

func NewTokener(secret string) JWTTokener {
	return JWTTokener{
		secret: secret,
	}
}

func (tokener JWTTokener) GetToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(tokenExpireTime).Unix(), // Время истечения
	}

	// Создаем токен с методом подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	return token.SignedString([]byte(tokener.secret))
}
func (tokener JWTTokener) GetUserId(token string) (int, error) {

	claims, err := tokener.getClaimsByToken(token)
	if err != nil {
		return 0, err
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		err := errors.New("user id not exist in claims")
		return 0, err
	}

	return int(userId), nil
}

func (tokener JWTTokener) getClaimsByToken(token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(tokener.secret), nil
	})

	var claims jwt.MapClaims
	if err != nil {
		return claims, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		return claims, errors.New("wrong claims format")
	}
	return claims, nil
}
