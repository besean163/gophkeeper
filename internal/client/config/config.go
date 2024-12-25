package config

const (
	DatabasePathDefault = "./data.db"
)

type Config struct {
	DatabasePath string
}

func NewConfig() Config {
	return Config{
		DatabasePath: DatabasePathDefault,
	}
}
