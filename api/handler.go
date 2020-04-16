package api

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/terotoi/marque/core"
	"github.com/jmoiron/sqlx"
)

// AnonAPIfunc a http.HandlerFunc with an error code.
type AnonAPIfunc func(http.ResponseWriter, *http.Request) error

// CreateAppHandler is a handler that calls the view specific function.
// AppHandler handles sessions, panic catching and error reporting.
// contentType:
//     ""     : no content type set
//     other  : literal content type
func createHandler(contentType string, cfg *core.Config, db *sqlx.DB) func(AnonAPIfunc) http.HandlerFunc {

	appHandler := func(next AnonAPIfunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Panic handler
			defer func() {
				if rec := recover(); rec != nil {
					if err, ok := rec.(error); ok {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusInternalServerError)
						json.NewEncoder(w).Encode(err)
					}
				}
			}()

			if contentType != "" {
				w.Header().Set("Content-Type", contentType)
			}

			err := next(w, r)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(false)

				//log.Printf("%s %s %s %s", r.RemoteAddr, r.Method, r.URL, err.Error())
				log.Print(err.Error())
				log.Print(string(debug.Stack()))
			}
		}
	}
	return appHandler
}
