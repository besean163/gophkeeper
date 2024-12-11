package jwttoken

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetJWTToken(secret string, claims jwt.MapClaims) (string, error) {

	// Создаем токен с методом подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	return token.SignedString([]byte(secret))
}

func GetClaimsByToken(secret, token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
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
