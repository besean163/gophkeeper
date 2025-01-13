package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type CardChanges struct {
	Created []*models.Card
	Updated []*models.Card
	Deleted []string
}

func NewCardChanges() CardChanges {
	return CardChanges{
		Created: make([]*models.Card, 0),
		Updated: make([]*models.Card, 0),
		Deleted: make([]string, 0),
	}
}
