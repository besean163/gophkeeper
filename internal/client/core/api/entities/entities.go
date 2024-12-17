package entities

type RegisterInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenOutput struct {
	Token string `json:"token"`
}
