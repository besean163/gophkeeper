package client

// Card модель карты клиента
type Card struct {
	UUID      string`gorm:"primaryKey"`
	UserID    int `json:"-"`
	Name      string
	Number    int
	Exp       string
	CVV       int
	CreatedAt int64 `gorm:"autoCreateTime:false"`
	UpdatedAt int64 `gorm:"autoUpdateTime:false"`
	DeletedAt int64 `gorm:"autoUpdateTime:false"`
	SyncedAt  int64 `gorm:"autoUpdateTime:false"`
}
