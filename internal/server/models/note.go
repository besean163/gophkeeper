package models

// Note модель заметки
type Note struct {
	ID        int    `json:"-"`
	UUID      string `json:"uuid"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:false"`
}
