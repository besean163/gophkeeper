package input

// CardCreate структура создания карты.
type CardCreate struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
	Exp    string `json:"exp"`
	CVV    int    `json:"cvv"`
}
