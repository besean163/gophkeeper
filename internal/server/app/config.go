package app

// Config структура конфигурации приложения
type Config struct {
	Host   string
	Secret string
}

// NewConfig создание структуры конфигурации приложения
func NewConfig() (*Config, error) {
	return &Config{
		Host:   "localhost:8080",
		Secret: "secret",
	}, nil
}
