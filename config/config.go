package config

import (
	"os"
)

type server struct {
	Port string
	Env  string
}

type database struct {
	URL string
}

type app struct {
	PublicKeyPath string
}

type Config struct {
	Server   server
	Database database
	App      app
}

func New() *Config {
	return &Config{
		Server: server{
			Port: os.Getenv("PORT"),
			Env:  os.Getenv("GO_ENV"),
		},
		Database: database{
			URL: os.Getenv("DATABASE_URL"),
		},
		App: app{
			PublicKeyPath: os.Getenv("PUBLIC_KEY_PATH"),
		},
	}
}
