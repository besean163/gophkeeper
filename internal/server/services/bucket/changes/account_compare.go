package changes

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type AccountCompare struct {
	Items        []*models.Account
	CompareItems []clientmodels.Account
}
