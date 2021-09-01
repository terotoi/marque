package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/utils"
)

// User account
type User struct {
	ID       int64
	Name     string
	Password string
	Admin    bool
}

// NewUser creates a new user object.
func NewUser(name, password string, admin bool, pwSecret string) (*User, error) {
	hashedPw, err := utils.HashPassword(pwSecret, password)
	if err != nil {
		return nil, err
	}

	u := User{
		Name:     name,
		Password: hashedPw,
		Admin:    admin,
	}
	return &u, nil
}

// Insert a user into the database.
func (u *User) Insert(tx *sqlx.Tx) error {
	query := "INSERT INTO users (name, password, admin) VALUES (?, ?, ?)"

	res, err := tx.Exec(query, u.Name, u.Password, u.Admin)
	if err != nil {
		return err
	}

	u.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update a user in the database.
func (u *User) Update(tx *sqlx.Tx) error {
	query := "UPDATE users SET name=?, password=?, admin=? WHERE id=?"
	_, err := tx.Exec(query, u.Name, u.Password, u.Admin, u.ID)
	if err != nil {
		log.Printf("User.Update: %s", err.Error())
	}
	return err
}

// UserByID returns a user object by id.
// Returns error if not found.
func UserByID(id int64, tx *sqlx.Tx) (*User, error) {
	u := User{}

	if err := tx.Get(&u, "SELECT * FROM users WHERE users.id=? LIMIT 1", id); err != nil {
		return nil, err
	}
	return &u, nil
}

// UserByName returns a user object by username.
func UserByName(name string, tx *sqlx.Tx) (*User, error) {
	u := []User{}

	if err := tx.Select(&u, "SELECT * FROM users WHERE users.name=? LIMIT 1", name); err != nil {
		return nil, err
	}

	if len(u) == 0 {
		return nil, nil
	} else {
		return &u[0], nil
	}
}

// NumUsers returns the number of users in the users table.
func NumUsers(tx *sqlx.Tx) (int, error) {
	var count int
	if err := tx.Get(&count, "SELECT COUNT(id) FROM users"); err != nil {
		return 0, err
	}

	return count, nil
}
