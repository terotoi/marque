package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
)

/**
 * Reports en error.
 *
 * Status code is sent to the browser, along with a JSON marshalled result and a message is written to the log.
 * Returns error only if the was an error doing this report.
 */
func report(status int, result interface{}, logMessage string, err error, stackTrace bool,
	w http.ResponseWriter, r *http.Request) error {

	header := fmt.Sprintf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

	txt := fmt.Sprintf("%s %s", header, logMessage)
	if err != nil {
		txt += ": " + err.Error()
	}
	log.Println(txt)

	if err != nil && stackTrace {
		lines := strings.Split(string(debug.Stack()), "\n")
		for _, line := range lines {
			log.Printf("%s %s", header, line)
		}
	}

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(result)
}

/**
 * Report an internal server error with a stack trace to the log.
 */
func reportInt(err error, w http.ResponseWriter, r *http.Request) error {
	return report(http.StatusInternalServerError, "Server error", "", err, true, w, r)
}

/**
 * Writes a JSON HTTP response body.
 */
func respondJSON(res interface{}, w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}
