package models

// ExternalAccount модель внешнего аккаунта для синхронизации
type ExternalAccount struct {
	UUID      string
	Name      string
	Login     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
	SyncedAt  int64
}
