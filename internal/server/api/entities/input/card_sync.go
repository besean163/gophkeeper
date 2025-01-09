package input

// CardSync структура карты для синхронизации карт.
type CardSync struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Exp       string `json:"exp"`
	CVV       int    `json:"cvv"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	SyncedAt  int64  `json:"synced_at"`
}
