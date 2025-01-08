package entities

// TokenOutput структура ответа успешной авторизации или регистрации.
type TokenOutput struct {
	Token string `json:"token"`
}
