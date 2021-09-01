package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/core"
	"github.com/terotoi/marque/models"
	"github.com/terotoi/marque/utils"
)

func normalizeBookmark(bm *models.Bookmark) error {
	bm.URL = strings.Trim(bm.URL, " \r\n\t")
	bm.Title = strings.Trim(bm.Title, " \r\n\t")
	if bm.Title == "" {
		title, err := utils.FetchTitle(bm.URL)
		if err != nil {
			return err
		}
		bm.Title = title
	}

	bm.UserID = nil // user.ID
	bm.Updated = time.Now()

	bm.CleanTags()
	return nil
}

func getBookmarksForUser(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {

		tx, err := db.Beginx()
		if err != nil {
			return reportInt(err, w, r)
		}
		defer tx.Rollback()

		user, err := userByAuth(r, site.JWTSecret, tx)
		if err != nil {
			return report(http.StatusForbidden, false, "JWT authentication failure", err, true, w, r)
		}

		bms, err := models.BookmarksAllForUser(user.ID, "ORDER BY ID desc", tx)
		if err != nil {
			return reportInt(err, w, r)
		}

		if err = tx.Commit(); err != nil {
			return reportInt(err, w, r)
		}

		return respondJSON(bms, w)
	}
}

func createBookmark(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		tx, err := db.BeginTxx(r.Context(), nil)
		if err != nil {
			return reportInt(err, w, r)
		}
		defer tx.Rollback()

		_, err = userByAuth(r, site.JWTSecret, tx)
		if err != nil {
			return report(http.StatusForbidden, false, "JWT authentication failure", err, true, w, r)
		}

		dec := json.NewDecoder(r.Body)

		var bm models.Bookmark
		err = dec.Decode(&bm)
		if err != nil {
			return reportInt(err, w, r)
		}

		if err := normalizeBookmark(&bm); err != nil {
			return reportInt(err, w, r)
		}

		// Check for duplicates.
		p, err := models.BookmarkByURL(bm.URL, tx)
		if err != nil {
			return reportInt(err, w, r)
		}

		if p != nil {
			return report(http.StatusOK, true, "bookmark already exists", nil, false, w, r)
		}

		if err := bm.Insert(tx); err != nil {
			return reportInt(err, w, r)
		}

		if err = tx.Commit(); err != nil {
			return reportInt(err, w, r)
		}

		return respondJSON(bm, w)
	}
}

func updateBookmark(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		tx, err := db.BeginTxx(r.Context(), nil)
		if err != nil {
			return reportInt(err, w, r)
		}
		defer tx.Rollback()

		_, err = userByAuth(r, site.JWTSecret, tx)
		if err != nil {
			return report(http.StatusForbidden, false, "JWT authentication failure", err, true, w, r)
		}

		dec := json.NewDecoder(r.Body)

		var bm models.Bookmark
		err = dec.Decode(&bm)
		if err != nil {
			return reportInt(err, w, r)
		}

		_, err = models.BookmarkByID(bm.ID, tx)
		if err != nil {
			return reportInt(err, w, r)
		}

		if err := normalizeBookmark(&bm); err != nil {
			return reportInt(err, w, r)
		}

		// Check for duplicates.
		p, err := models.BookmarkByURL(bm.URL, tx)
		if err != nil {
			return reportInt(err, w, r)
		}

		if p != nil && p.ID != bm.ID {
			return report(http.StatusOK, true, "bookmark already exists", nil, false, w, r)
		}

		if err := bm.Update(tx); err != nil {
			return reportInt(err, w, r)
		}

		if err = tx.Commit(); err != nil {
			return reportInt(err, w, r)
		}

		return respondJSON(bm, w)
	}
}

func deleteBookmark(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		tx, err := db.BeginTxx(r.Context(), nil)
		if err != nil {
			return reportInt(err, w, r)
		}
		defer tx.Rollback()

		_, err = userByAuth(r, site.JWTSecret, tx)
		if err != nil {
			return report(http.StatusForbidden, false, "JWT authentication failure", err, true, w, r)
		}

		id, err := strconv.ParseInt(r.URL.Path[len("/api/bookmark/delete/"):], 10, 64)
		if err != nil {
			return reportInt(err, w, r)
		}

		bm, err := models.BookmarkByID(id, tx)
		if err != nil {
			return reportInt(err, w, r)
		}

		if err := bm.Delete(tx); err != nil {
			return reportInt(err, w, r)
		}

		if err = tx.Commit(); err != nil {
			return reportInt(err, w, r)
		}

		return respondJSON(true, w)
	}
}
