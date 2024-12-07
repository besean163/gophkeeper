package models

type User struct {
	ID    int
	Login string
}

type Test struct {
	Key string `json:"key"`
}
