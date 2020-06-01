package core

import (
	"log"

	"github.com/jmoiron/sqlx"

	//_ "github.com/lib/pq"           // For PostgreSQL driver
	_ "github.com/mattn/go-sqlite3" // SQLite
)

// SetupDB connects to a database and initializes it, if needed.
func SetupDB(cfg *Config) (*sqlx.DB, error) {
	var db *sqlx.DB

	log.Printf("Connecting to %s %s\n", cfg.DatabaseType, cfg.DatabaseString)

	db, err := sqlx.Connect(cfg.DatabaseType, cfg.DatabaseString)
	if err != nil {
		return nil, err
	}

	if err := createDatabaseSqlite(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabaseSqlite(db *sqlx.DB) error {
	query := "CREATE TABLE IF NOT EXISTS bookmarks " +
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER, title VARCHAR NOT NULL, " +
		"url VARCHAR NOT NULL, tags VARCHAR NOT NULL DEFAULT '', " +
		"notes VARCHAR NOT NULL DEFAULT '', " +
		"updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)"
	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
