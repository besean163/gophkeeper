package entities

import (
	"time"

	"github.com/besean163/gophkeeper/internal/server/models"
)

type GetTokenInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenOutput struct {
	Token string `json:"token"`
}

type AccountsOutput struct {
	Accounts []*models.Account `json:"accounts"`
}

type AccountOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
