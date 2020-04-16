package main

import (
	"net/http"

	"github.com/terotoi/marque/api"
	"github.com/terotoi/marque/core"
	"github.com/jmoiron/sqlx"
	"github.com/markbates/pkger"
)

func setupRoutes(site *core.Site, db *sqlx.DB) {
	api.SetupRoutes(site.Config, site, db)

	http.Handle("/", //http.StripPrefix("/static/",
		http.FileServer(pkger.Dir("/public")))
}
