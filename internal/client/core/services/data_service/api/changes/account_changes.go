package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

type AccountChanges struct {
	Created []models.Account
	Updated []models.Account
	Deleted []models.Account
}

func NewAccountChanges() AccountChanges {
	return AccountChanges{
		Created: make([]models.Account, 0),
		Updated: make([]models.Account, 0),
		Deleted: make([]models.Account, 0),
	}
}
