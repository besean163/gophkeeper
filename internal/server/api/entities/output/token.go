package output

// Token структура ответа успешной авторизации или регистрации.
type Token struct {
	Token string `json:"token"`
}
