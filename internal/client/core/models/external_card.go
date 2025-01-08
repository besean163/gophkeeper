package models

// ExternalCard модель входящей карты для синхронизации
type ExternalCard struct {
	ID        int
	UUID      string
	UserID    int
	Name      string
	Number    int
	Exp       string
	CVV       int
	CreatedAt int64
	UpdatedAt int64
}
