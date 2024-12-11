package models

type User struct {
	ID       int
	Login    string
	Password string
	Token    string
}

type Test struct {
	Key string `json:"key"`
}
