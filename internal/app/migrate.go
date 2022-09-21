package app

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	type dbConfig struct {
		Postgres struct {
			Connection string `envconfig:"DATABASE_URL"`
		}
	}

	cfg := &dbConfig{}

	err := envconfig.Process("envconfig", cfg)
	if err != nil {
		log.Fatalf("migrate: environment variable not declared: DATABASE_URL")
	}

	m, err := migrate.New("file://migrations", cfg.Postgres.Connection)
	if err != nil {
		log.Fatalf("migrate: postgres connection")
	}

	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: up error: %s", err)
		return
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate: no changes")
		return
	}

	log.Printf("migrate: up success")
}
