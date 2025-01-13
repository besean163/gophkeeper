// Package dependencies предоставляет зависимости для работы маршрутов сервера
package dependencies

import (
	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	apitoken "github.com/besean163/gophkeeper/internal/utils/api_token"
)

// Dependencies структура со списком зависимостей.
type Dependencies struct {
	Logger        logger.Logger
	AuthService   interfaces.AuthService
	BucketService interfaces.BucketService
	Tokener       apitoken.Tokener
}
