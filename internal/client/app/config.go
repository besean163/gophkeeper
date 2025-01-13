package app

const (
	DatabasePathDefault = "./data.db"
)

// Config структура конфигурации приложения
type Config struct {
	DatabasePath string
}

// NewConfig создание структуры конфигурации приложения
func NewConfig() Config {
	return Config{
		DatabasePath: DatabasePathDefault,
	}
}
