package models

// ExternalCard модель внешней карты для синхронизации
type ExternalCard struct {
	UUID      string
	Name      string
	Number    int
	Exp       string
	CVV       int
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64
	SyncedAt  int64
}
