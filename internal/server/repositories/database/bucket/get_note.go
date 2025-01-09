package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetNote получение заметки
func (r Repository) GetNote(uuid string) (*models.Note, error) {
	var item *models.Note
	r.db.Where("uuid = ?", uuid).Find(&item)
	if item.UUID == "" {
		return nil, nil
	}
	return item, nil
}
