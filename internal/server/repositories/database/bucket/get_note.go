package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// GetNote получение заметки
func (r Repository) GetNote(uuid string) (*models.Note, error) {
	var item *models.Note
	r.db.Where("uuid = ?", uuid).Find(&item)
	if item.ID == 0 {
		return nil, nil
	}
	return item, nil
}
