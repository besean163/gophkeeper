package server

// User модель пользователя сервера
type User struct {
	ID        int
	UUID      string
	Login     string
	Password  string
	CreatedAt int64
}
