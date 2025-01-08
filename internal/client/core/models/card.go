package models

// Card модель карты
type Card struct {
	ID        int
	UUID      string
	UserID    int
	Name      string
	Number    int
	Exp       string
	CVV       int
	CreatedAt int64 `gorm:"autoCreateTime:false"`
	UpdatedAt int64 `gorm:"autoUpdateTime:false"`
	DeletedAt int64 `gorm:"autoUpdateTime:false"`
	SyncedAt  int64 `gorm:"autoUpdateTime:false"`
}
