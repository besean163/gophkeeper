package entities

// NoteSyncInput структура заметки для синхронизации заметок.
type NoteSyncInput struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
	SyncedAt  int64  `json:"synced_at"`
}
