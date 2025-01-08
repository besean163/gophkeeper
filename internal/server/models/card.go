package models

// Card модель карты
type Card struct {
	ID        int    `json:"-"`
	UUID      string `json:"uuid"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Exp       string `json:"exp"`
	CVV       int    `json:"cvv"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:false"`
}
