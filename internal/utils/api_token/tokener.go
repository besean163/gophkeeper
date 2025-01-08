package apitoken

import (
	"github.com/besean163/gophkeeper/internal/server/models"
)

type Tokener interface {
	GetToken(user *models.User) (string, error)
	GetUserId(token string) (int, error)
}
