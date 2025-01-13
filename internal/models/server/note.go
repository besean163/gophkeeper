package server

// Note модель заметки сервера
type Note struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:false"`
}
