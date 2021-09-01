package api

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/core"
)

// SetupRoutes sets up the HTTP handlers.
func SetupRoutes(cfg *core.Config, site *core.Site, db *sqlx.DB) {
	handlerJSON := createHandler("application/json", cfg, db)

	http.HandleFunc("/api/login", handlerJSON(login(site, db)))
	http.HandleFunc("/api/get_bookmarks", handlerJSON(getBookmarksForUser(site, db)))
	http.HandleFunc("/api/bookmark/create", handlerJSON(createBookmark(site, db)))
	http.HandleFunc("/api/bookmark/update", handlerJSON(updateBookmark(site, db)))
	http.HandleFunc("/api/bookmark/delete/", handlerJSON(deleteBookmark(site, db)))
	http.HandleFunc("/api/user/set_password", handlerJSON(setPassword(site, db)))
}
