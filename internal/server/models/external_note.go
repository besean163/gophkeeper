package models

// ExternalNote модель внешней заметки для синхронизации
type ExternalNote struct {
	UUID      string
	Name      string
	Content   string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
	SyncedAt  int64
}
