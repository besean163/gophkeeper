package entities

// NotesSyncInput структура для синхронизации заметок.
type NotesSyncInput struct {
	Notes []NoteSyncInput `json:"notes"`
}
