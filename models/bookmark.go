package models

import (
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// Bookmark datatype
type Bookmark struct {
	ID       int64
	UserID   *int64 `db:"user_id"`
	Title    string
	URL      string
	Tags     []string `db:"-"`
	TagsText string   `db:"tags" json:"-"` // Tags converted to string, for storing into sqlite.
	Notes    string
	Updated  time.Time
}

const bookmarkInsertNames = "(user_id, title, url, tags, notes, updated)"
const bookmarkInsertParams = "(?, ?, ?, ?, ?, ?)"
const bookmarkUpdateParams = "user_id=?, title=?, url=?, tags=?, notes=?, updated=? "
const bookmarkSelectFields = "bookmarks.id, bookmarks.user_id, bookmarks.title, " +
	"bookmarks.url, bookmarks.tags, bookmarks.notes, bookmarks.updated "

// NewBookmark creates a new bookmark with ID -1.
func NewBookmark(title, url, notes string, tags []string) *Bookmark {
	bm := Bookmark{
		ID:       -1,
		UserID:   nil,
		Title:    title,
		URL:      url,
		Tags:     tags,
		TagsText: tagsToDB(tags),
		Notes:    notes,
		Updated:  time.Now(),
	}
	return &bm
}

func tagsToDB(tags []string) string {
	return strings.Join(tags, " ")
}

func tagsFromDB(text string) []string {
	if text == "" {
		return nil
	}
	return strings.Split(text, " ")
}

// CleanTags cleans up the tags to a list separated by commas.
func (bm *Bookmark) CleanTags() {
	var r []string

	for _, tag := range bm.Tags {
		tag = strings.ToLower(tag)

		if len(tag) > 0 {
			if tag[0] != '#' {
				tag = "#" + tag
				r = append(r, tag)
			} else if len(tag) > 1 {
				r = append(r, tag)
			}
		}
	}

	bm.Tags = r
	bm.TagsText = tagsToDB(r)
}

// Insert a bookmark into the database.
func (bm *Bookmark) Insert(db *sqlx.DB) error {
	query := "INSERT INTO bookmarks " + bookmarkInsertNames + " VALUES " +
		bookmarkInsertParams

	res, err := db.Exec(query, bm.UserID, bm.Title, bm.URL,
		bm.TagsText, bm.Notes, bm.Updated)
	if err != nil {
		return err
	}

	bm.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update a bookmark in the database.
func (bm *Bookmark) Update(db *sqlx.DB) error {
	query := "UPDATE bookmarks SET " + bookmarkUpdateParams + " WHERE id=?"
	_, err := db.Exec(query, bm.UserID, bm.Title, bm.URL, bm.TagsText,
		bm.Notes, bm.Updated, bm.ID)
	//affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	return err
}

// Delete deletes a bookmark from the database.
func (bm *Bookmark) Delete(db *sqlx.DB) error {
	query := "DELETE FROM bookmarks WHERE id=?"
	if _, err := db.Exec(query, bm.ID); err != nil {
		log.Printf("Bookmark.Delete: %s", err.Error())
		return err
	}

	return nil
}

// BookmarkByID returns a Bookmark by ID.
func BookmarkByID(id int64, db *sqlx.DB) (*Bookmark, error) {
	bm := Bookmark{}

	query := "SELECT " + bookmarkSelectFields + " FROM bookmarks " +
		" WHERE bookmarks.id=? LIMIT 1"

	if err := db.Get(&bm, query, id); err != nil {
		return nil, err
	}
	bm.Tags = tagsFromDB(bm.TagsText)

	return &bm, nil
}

// BookmarkByURL returns all bookmark with matching URL, or nil.
func BookmarkByURL(url string, db *sqlx.DB) (*Bookmark, error) {
	bms := []Bookmark{}

	query := "SELECT " + bookmarkSelectFields + " FROM bookmarks " +
		" WHERE bookmarks.url=? AND user_id IS NULL LIMIT 1"

	if err := db.Select(&bms, query, url); err != nil {
		return nil, err
	}

	if len(bms) > 0 {
		bms[0].Tags = tagsFromDB(bms[0].TagsText)
		return &bms[0], nil
	}
	return nil, nil
}

// BookmarksAll returns all bookmarks.
func BookmarksAll(orderBy string, db *sqlx.DB) ([]*Bookmark, error) {
	var bms []*Bookmark

	if err := db.Select(&bms, "SELECT "+bookmarkSelectFields+
		" FROM bookmarks "+orderBy); err != nil {
		return nil, err
	}

	for _, bm := range bms {
		bm.Tags = tagsFromDB(bm.TagsText)
	}
	return bms, nil
}

// BookmarksAllForUser returns all bookmarks owned by a user or not owned
// by anone (user is NULL).
func BookmarksAllForUser(userID int64, orderBy string, db *sqlx.DB) ([]*Bookmark, error) {
	var bms []*Bookmark

	if err := db.Select(&bms, "SELECT "+bookmarkSelectFields+
		" FROM bookmarks WHERE user_id=? OR user_id IS NULL "+orderBy, userID); err != nil {
		return nil, err
	}

	for _, bm := range bms {
		bm.Tags = tagsFromDB(bm.TagsText)
	}

	return bms, nil
}
