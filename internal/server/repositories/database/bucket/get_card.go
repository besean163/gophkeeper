package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetCard получение карты
func (r Repository) GetCard(uuid string) (*models.Card, error) {
	var item *models.Card
	r.db.Where("uuid = ?", uuid).Find(&item)
	if item.UUID == "" {
		return nil, nil
	}
	return item, nil
}
