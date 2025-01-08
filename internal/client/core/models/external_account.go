package models

// ExternalAccount модель входящего аккаунта для синхронизации 
type ExternalAccount struct {
	ID        int
	UUID      string
	UserID    int
	Name      string
	Login     string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}
