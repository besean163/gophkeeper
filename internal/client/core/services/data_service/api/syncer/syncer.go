package syncer

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

type Syncer interface {
	Sync(user models.User, entities ...int) error
}
