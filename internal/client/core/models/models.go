package models

type User struct {
	ID        int
	Login     string
	Password  string
	Token     string
	CreatedAt int
}

type Account struct {
	ID        int
	UserID    int
	Name      string
	Login     string
	Password  string
	CreatedAt int
	UpdatedAt int
	DeletedAt int
	SyncedAt  int
}
