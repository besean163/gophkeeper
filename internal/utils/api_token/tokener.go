package apitoken

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type Tokener interface {
	GetToken(user *models.User) (string, error)
	GetUserId(token string) (int, error)
}
