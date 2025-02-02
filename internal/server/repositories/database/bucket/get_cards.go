package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetCards получение списка карт
func (r Repository) GetCards(user models.User) ([]*models.Card, error) {
	items := []*models.Card{}
	r.db.Where("user_id = ?", user.ID).Find(&items)
	return items, nil
}
