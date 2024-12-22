package models

import "time"

type Account struct {
	ID        int    `json:"id"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (Account) TableName() string {
// 	return "accounts"
// }
