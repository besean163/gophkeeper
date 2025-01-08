package models

// ExternalNote модель входящей заметки для синхронизации
type ExternalNote struct {
	ID        int
	UUID      string
	UserID    int
	Name      string
	Content   string
	CreatedAt int64
	UpdatedAt int64
}
