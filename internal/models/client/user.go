package client

// User модель пользователя клиента
type User struct {
	ID        int
	Login     string
	Password  string
	Token     string
	CreatedAt int64
}
