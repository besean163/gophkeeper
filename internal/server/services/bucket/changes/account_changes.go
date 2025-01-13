package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type AccountChanges struct {
	Created []*models.Account
	Updated []*models.Account
	Deleted []string
}

func NewAccountChanges() AccountChanges {
	return AccountChanges{
		Created: make([]*models.Account, 0),
		Updated: make([]*models.Account, 0),
		Deleted: make([]string, 0),
	}
}
