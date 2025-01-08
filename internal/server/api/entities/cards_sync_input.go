package entities

// CardsSyncInput структура для синхронизяции карт.
type CardsSyncInput struct {
	Cards []CardSyncInput `json:"cards"`
}
