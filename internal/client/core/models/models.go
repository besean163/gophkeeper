package models

type User struct {
	ID        int
	Login     string
	Password  string
	Token     string
	CreatedAt int
}
