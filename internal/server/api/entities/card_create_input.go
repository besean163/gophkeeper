package entities

// CardCreateInput структура создания карты.
type CardCreateInput struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
	Exp    string `json:"exp"`
	CVV    int    `json:"cvv"`
}
