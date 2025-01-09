package client

// Note модель заметки клиента
type Note struct {
	UUID      string`gorm:"primaryKey"`
	UserID    int `json:"-"`
	Name      string
	Content   string
	CreatedAt int64 `gorm:"autoCreateTime:false"`
	UpdatedAt int64 `gorm:"autoUpdateTime:false"`
	DeletedAt int64 `gorm:"autoUpdateTime:false"`
	SyncedAt  int64 `gorm:"autoUpdateTime:false"`
}
