package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

type CardChanges struct {
	Created []models.Card
	Updated []models.Card
	Deleted []models.Card
}

func NewCardChanges() CardChanges {
	return CardChanges{
		Created: make([]models.Card, 0),
		Updated: make([]models.Card, 0),
		Deleted: make([]models.Card, 0),
	}
}
