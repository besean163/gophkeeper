package models

import "time"

type Account struct {
	ID        int       `json:"id"`
	UserID    int       `json:"-"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// func (Account) TableName() string {
// 	return "accounts"
// }
