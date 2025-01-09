package input

// AccountCreate структура для создания аккаунта.
type AccountCreate struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
