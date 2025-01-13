package server

// Card модель карты  сервера
type Card struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	UserID    int    `json:"-"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Exp       string `json:"exp"`
	CVV       int    `json:"cvv"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:false"`
}
