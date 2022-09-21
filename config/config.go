package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Name        string `envconfig:"APP_NAME"`
		Port        string `envconfig:"APP_PORT"`
		Environment string `envconfig:"APP_ENV"`
		Version     string `envconfig:"APP_VERSION"`
	}
	Postgres struct {
		Connection string `envconfig:"DATABASE_URL"`
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("envconfig", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.App.Environment != "development" {
		cfg.App.Port = os.Getenv("PORT")
	}

	return cfg, nil
}
