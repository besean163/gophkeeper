package ctxuser

import (
	"context"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
)

func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(entities.RequestUserKey("user")).(*models.User)
	return user, ok
}
