package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// GetCard получение карты
func (r Repository) GetCard(uuid string) (*models.Card, error) {
	var item *models.Card
	r.db.Where("uuid = ?", uuid).Find(&item)
	if item.ID == 0 {
		return nil, nil
	}
	return item, nil
}
