package app

import (
	"fmt"
	"log"

	"github.com/alfonsovgs/venture/config"
	"github.com/alfonsovgs/venture/internal/controller/http"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func Run(cfg *config.Config) {
	url, err := buildPostgresConnectionString(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - buildPostgresConnectionString: %w", err))
	}

	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - sqlx.Connect: %w", err))
	}

	dependencies := http.NewContainer(db)

	server := http.NewServer(dependencies)
	server.Middlewares(http.WithLogger(), http.WithRecover(), http.WithCORS(), http.WithGzip())
	server.MapRoutes()
	server.Start(cfg.App.Port)
}

func buildPostgresConnectionString(cfg *config.Config) (string, error) {
	return pq.ParseURL(cfg.Postgres.Connection)
}
