package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/api"
	"github.com/terotoi/marque/core"
)

func setupRoutes(site *core.Site, db *sqlx.DB) {
	api.SetupRoutes(site.Config, site, db)

	http.Handle("/", http.FileServer(http.Dir(site.Config.PublicDir)))
}
