package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

type AccountCompare struct {
	Items        []models.Account
	CompareItems []servermodels.Account
}
