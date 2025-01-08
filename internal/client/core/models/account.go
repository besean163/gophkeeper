package models

// Account модель ядра
type Account struct {
	ID        int
	UUID      string
	UserID    int
	Name      string
	Login     string
	Password  string
	CreatedAt int64 `gorm:"autoCreateTime:false"`
	UpdatedAt int64 `gorm:"autoUpdateTime:false"`
	DeletedAt int64 `gorm:"autoUpdateTime:false"`
	SyncedAt  int64 `gorm:"autoUpdateTime:false"`
}
