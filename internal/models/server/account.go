package server

// Account модель аккаунта сервера
type Account struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:false"`
}
