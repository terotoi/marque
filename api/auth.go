package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/jwt"
	"github.com/terotoi/marque/models"
)

type jwtAuth struct {
	Subject string `json:"sub"`
}

// Decode a JWT authorizatoin payload and uses the "sub" claim
// to find a correspong user object.
// If user is not found returns nil, "user not found" error
func userByToken(payload []byte, db *sqlx.Tx) (*models.User, error) {
	var auth jwtAuth

	if err := json.Unmarshal(payload, &auth); err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(auth.Subject)
	if err != nil {
		return nil, err
	}

	user, err := models.UserByID(int64(id), db)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Decode HTTP Authorization header with Bearer JWT Token.
// See userByToken for more details.
func userByAuth(r *http.Request, secret []byte, tx *sqlx.Tx) (*models.User, error) {
	payload, err := jwt.ParseAuthorization(r.Header.Get("Authorization"), secret)
	if err != nil {
		return nil, err
	}

	return userByToken(payload, tx)
}
