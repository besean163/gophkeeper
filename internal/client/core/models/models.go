package models

import "time"

type User struct {
	ID        int
	Login     string
	Password  string
	Token     string
	CreatedAt time.Time
}

type Account struct {
	ID        int
	UserID    int
	Name      string
	Login     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	SyncedAt  time.Time
}
