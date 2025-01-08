package models

// User модель пользователя
type User struct {
	ID        int
	UUID      string
	Login     string
	Password  string
	CreatedAt int64
}
