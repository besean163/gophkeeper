package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// GetNotes получение списка заметок
func (r Repository) GetNotes(user models.User) ([]*models.Note, error) {
	items := []*models.Note{}
	r.db.Where("user_id = ?", user.ID).Find(&items)
	return items, nil
}
