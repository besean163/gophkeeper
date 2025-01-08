package models

// Account модель аккаунта
type Account struct {
	ID        int    `json:"-"`
	UUID      string `json:"uuid"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:false"`
}
