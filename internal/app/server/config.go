package server

type Config struct {
	Host string
}

func ReadConfig() (*Config, error) {
	return &Config{
		Host: "localhost:8080",
	}, nil
}
