package main

import (
	"log"
	"net/http"

	"github.com/terotoi/marque/core"
	"github.com/jmoiron/sqlx"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func serve(site *core.Site, db *sqlx.DB) error {
	setupRoutes(site, db)

	log.Printf("Listening on %s\n", site.Config.ListenAddress)
	return http.ListenAndServe(site.Config.ListenAddress, logRequest(http.DefaultServeMux))
}
