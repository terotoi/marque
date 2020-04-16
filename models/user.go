package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// User account
type User struct {
	ID       int64
	Name     string
	Password *string
}

// NewUser creates a new user object.
func NewUser(name string) *User {
	u := User{
		Name: name,
	}
	return &u
}

const usersInsertNames = "(name, password)"
const usersInsertParams = "($1, $2)"
const usersUpdateParams = "name=$2, password=$3"
const usersSelectFields = "users.id, users.name, users.password"

// Insert a user into the database.
func (u *User) Insert(db *sqlx.DB) (int64, error) {
	query := "INSERT INTO users " + usersInsertNames + " VALUES " +
		usersInsertParams +
		" RETURNING id"

	// sqlite3: datetime($5)

	row := db.QueryRow(query, u.Name, u.Password)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		log.Printf("User.Insert: %s", err.Error())
	}
	u.ID = id
	return id, err
}

// Update a user in the database.
func (u *User) Update(db *sqlx.DB) error {
	query := "UPDATE users SET " + usersUpdateParams + " WHERE id=$1"
	_, err := db.Exec(query, u.ID, u.Name, u.Password)
	if err != nil {
		log.Printf("User.Update: %s", err.Error())
	}
	return err
}

// UserByID returns a user object by id.
// Returns error if not found.
func UserByID(id int64, db *sqlx.DB) (*User, error) {
	u := User{}

	if err := db.Get(&u, "SELECT "+usersSelectFields+" FROM users "+
		"WHERE users.id=$1", id); err != nil {
		return nil, err
	}
	return &u, nil
}

// UserByName returns a user object by username.
func UserByName(name string, db *sqlx.DB) (*User, error) {
	u := []*User{}

	if err := db.Select(&u, "SELECT "+usersSelectFields+" FROM users "+
		"WHERE users.name=$1", name); err != nil {
		return nil, err
	}

	if len(u) > 0 {
		return u[0], nil
	}
	return nil, nil
}
