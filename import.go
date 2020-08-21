package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/models"
)

// ImportJSON imports bookmarks from JSON file with a custom format.
func ImportJSON(filename string, db *sqlx.DB) error {
	fmt.Printf("Importing bookmarks from %s\n", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var bms []models.Bookmark
	if err := json.Unmarshal(data, &bms); err != nil {
		return err
	}

	for _, bm := range bms {
		bm.CleanTags()
		if err := bm.Insert(db); err != nil {
			return err
		}
	}

	fmt.Printf("Imported %d bookmarks.\n", len(bms))
	return nil
}

// ExportJSON exports bookmarks to JSON file with a custom format.
func ExportJSON(filename string, db *sqlx.DB) error {
	bms, err := models.BookmarksAll("", db)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(bms, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, data, 0600); err != nil {
		return err
	}

	fmt.Printf("Exported %d bookmarks.\n", len(bms))
	return nil
}
