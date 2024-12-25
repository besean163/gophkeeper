package app

type Config struct {
	Host   string
	Secret string
}

func ReadConfig() (*Config, error) {
	return &Config{
		Host:   "localhost:8080",
		Secret: "secret",
	}, nil
}
