package ctxuser

import (
	"context"

	"github.com/besean163/gophkeeper/internal/server/api/entities"
	"github.com/besean163/gophkeeper/internal/server/models"
)

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(entities.RequestUserKey("user")).(*models.User)
	return user, ok
}
