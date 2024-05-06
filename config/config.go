package config

import (
	"github.com/dscamargo/go_app_template/pkg"
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
			Port: pkg.GetEnvOrDefault("PORT", "8080"),
			Env:  pkg.GetEnvOrDefault("GO_ENV", "development"),
		},
		Database: database{
			URL: pkg.GetEnvOrDefault("DATABASE_URL", ""),
		},
		App: app{
			PublicKeyPath: pkg.GetEnvOrDefault("PUBLIC_KEY_PATH", "ssl/public.key"),
		},
	}
}
