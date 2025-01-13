package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

type CardCompare struct {
	Items        []models.Card
	CompareItems []servermodels.Card
}
