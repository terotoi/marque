package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/terotoi/marque/core"
	"github.com/terotoi/marque/models"
	"github.com/terotoi/marque/utils"
	"github.com/jmoiron/sqlx"
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
		/*user, err := userByAuth(r, site.JWTSecret, db)
		if err != nil {
			return err
		}*/

		/*
			id, err := strconv.ParseInt(r.URL.Path[len("/api/bookmark/get"):], 10, 64)
			if err != nil {
				return err
			}*/

		bms, err := models.BookmarksAllForUser(-1, "ORDER BY ID desc", db)
		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bms)
		return nil
	}
}

func createBookmark(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		/*user, err := userByAuth(r, site.JWTSecret, db)
		if err != nil {
			return err
		}*/

		dec := json.NewDecoder(r.Body)

		var bm models.Bookmark
		err := dec.Decode(&bm)
		if err != nil {
			return err
		}

		if err := normalizeBookmark(&bm); err != nil {
			return err
		}

		// Check for duplicates.
		p, err := models.BookmarkByURL(bm.URL, db)
		if err != nil {
			return err
		}

		if p != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(true)
			return nil
		}

		if err := bm.Insert(db); err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bm)
		return nil
	}
}

func updateBookmark(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		/*user, err := userByAuth(r, site.JWTSecret, db)
		if err != nil {
			return err
		}*/
		dec := json.NewDecoder(r.Body)

		var bm models.Bookmark
		err := dec.Decode(&bm)
		if err != nil {
			return err
		}

		_, err = models.BookmarkByID(bm.ID, db)
		if err != nil {
			return err
		}

		if err := normalizeBookmark(&bm); err != nil {
			return err
		}

		// Check for duplicates.
		p, err := models.BookmarkByURL(bm.URL, db)
		if err != nil {
			return err
		}

		if p != nil && p.ID != bm.ID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(true)
			return nil
		}

		if err := bm.Update(db); err != nil {
			return err
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bm)
		return nil
	}
}

func deleteBookmark(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		/*user, err := userByAuth(r, site.JWTSecret, db)
		if err != nil {
			return err
		}*/

		id, err := strconv.ParseInt(r.URL.Path[len("/api/bookmark/delete/"):], 10, 64)
		if err != nil {
			return err
		}

		bm, err := models.BookmarkByID(id, db)
		if err != nil {
			return err
		}

		if err := bm.Delete(db); err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
		return nil
	}
}
