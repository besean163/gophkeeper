package models

// User модель пользователя
type User struct {
	ID        int
	Login     string
	Password  string
	Token     string
	CreatedAt int64
}
