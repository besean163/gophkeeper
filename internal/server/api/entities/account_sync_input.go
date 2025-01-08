package entities

// AccountSyncInput структура аккаунта для синхронизации аккаунтов.
type AccountSyncInput struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	SyncedAt  int64  `json:"synced_at"`
}
