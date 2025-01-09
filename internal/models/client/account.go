package client

// Account модель аккаунта клиента
type Account struct {
	UUID      string `gorm:"primaryKey"`
	UserID    int    `json:"-"`
	Name      string
	Login     string
	Password  string
	CreatedAt int64 `gorm:"autoCreateTime:false"`
	UpdatedAt int64 `gorm:"autoUpdateTime:false"`
	DeletedAt int64 `gorm:"autoUpdateTime:false"`
	SyncedAt  int64 `gorm:"autoUpdateTime:false"`
}
