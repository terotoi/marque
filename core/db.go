package core

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/models"

	//_ "github.com/lib/pq"           // For PostgreSQL driver
	_ "github.com/mattn/go-sqlite3" // SQLite
)

// SetupDB connects to a database and initializes it, if needed.
func SetupDB(cfg *Config) (*sqlx.DB, error) {
	var db *sqlx.DB

	log.Printf("Connecting to %s %s\n", cfg.DatabaseType, cfg.DatabaseConfig)

	db, err := sqlx.Connect(cfg.DatabaseType, cfg.DatabaseConfig)
	if err != nil {
		return nil, err
	}

	if err := createDatabaseSqlite(context.Background(), cfg, db); err != nil {
		return nil, err
	}
	return db, nil
}

func createDatabaseSqlite(ctx context.Context, cfg *Config, db *sqlx.DB) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	createBookmarks := "CREATE TABLE IF NOT EXISTS bookmarks " +
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"user_id INTEGER, title VARCHAR NOT NULL, " +
		"url VARCHAR NOT NULL, tags VARCHAR NOT NULL DEFAULT '', " +
		"notes VARCHAR NOT NULL DEFAULT '', " +
		"updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)"
	if _, err := tx.Exec(createBookmarks); err != nil {
		return err
	}

	createUsers := "CREATE TABLE IF NOT EXISTS users " +
		"(id INTEGER PRIMARY KEY AUTOINCREMENT, " +
		"name VARCHAR UNIQUE NOT NULL, " +
		"password VARCHAR NOT NULL, " +
		"admin BOOL NOT NULL DEFAULT false)"
	if _, err := tx.Exec(createUsers); err != nil {
		return err
	}

	if cfg.InitialUser != "" && cfg.InitialPassword != "" {
		numUsers, err := models.NumUsers(tx)
		if err != nil {
			return err
		}

		if numUsers == 0 {
			log.Printf("Creating an initial admin user: %s", cfg.InitialUser)

			user, err := models.NewUser(cfg.InitialUser, cfg.InitialPassword, true, cfg.PasswordSecret)
			if err != nil {
				return err
			}

			if err := user.Insert(tx); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}
