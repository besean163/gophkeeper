package changes

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type CardCompare struct {
	Items        []*models.Card
	CompareItems []clientmodels.Card
}
