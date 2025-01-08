package entities

// AccountSyncInput структура для синхронизации аккаунтов.
type AccountsSyncInput struct {
	Accounts []AccountSyncInput `json:"accounts"`
}
