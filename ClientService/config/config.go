package config

type Config struct {
	Port string
	Host string
}

func NewConfig() *Config {
	return &Config{
		Port: "8070",
		Host: "127.0.0.1",
	}
}
