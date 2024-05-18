package config

import "os"

type Config struct {
	Port string
}

func LoadConfig() *Config {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}

	return &Config{
		Port: port,
	}
}
