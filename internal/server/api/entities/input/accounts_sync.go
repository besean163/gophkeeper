package input

// AccountSyncInput структура для синхронизации аккаунтов.
type AccountsSync struct {
	Accounts []AccountSync `json:"accounts"`
}
