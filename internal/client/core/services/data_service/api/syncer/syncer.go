package syncer

import "github.com/besean163/gophkeeper/internal/client/core/models"

type Syncer interface {
	SyncAll(user models.User) error
	SyncAccounts(user models.User) error
	SyncNotes(user models.User) error
	SyncCards(user models.User) error
}
