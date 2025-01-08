package entities

// AccountPostInput структура для создания аккаунта.
type AccountPostInput struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
