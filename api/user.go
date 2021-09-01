package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/terotoi/marque/core"
	"github.com/terotoi/marque/jwt"
	"github.com/terotoi/marque/models"
	"github.com/terotoi/marque/utils"
)

// LoginRequest is used for loggin in and getting the jwt auth token
type LoginRequest struct {
	Username string
	Password string
}

// SetPasswordRequest is used for changing passwords.
type SetPasswordRequest struct {
	Username    string
	OldPassword string
	NewPassword string
}

func login(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		dec := json.NewDecoder(r.Body)

		var loginReq LoginRequest
		err := dec.Decode(&loginReq)
		if err != nil {
			return err
		}

		tx, err := db.BeginTxx(r.Context(), nil)
		if err != nil {
			return err
		}
		defer tx.Rollback()

		user, err := models.UserByName(loginReq.Username, tx)
		if err != nil {
			return err
		}

		if user == nil {
			return report(http.StatusForbidden, false,
				fmt.Sprintf("User %s not found", loginReq.Username),
				err, false, w, r)
		}

		if !utils.ComparePassword(site.Config.PasswordSecret, loginReq.Password,
			user.Password) {
			return report(http.StatusForbidden, false,
				fmt.Sprintf("Password mismatch for %s", loginReq.Username),
				err, false, w, r)
		}

		token, err := jwt.CreateToken(site.JWTSecret,
			jwtAuth{
				strconv.Itoa(int(user.ID)),
			})

		if err != nil {
			return err
		}

		response := struct {
			AuthToken string
			Username  string
			IsAdmin   bool
		}{
			token,
			user.Name,
			user.Admin,
		}

		if err := tx.Commit(); err != nil {
			return reportInt(err, w, r)
		}
		return respondJSON(response, w)
	}
}

// SetPassword changes a password. Non-admins can only change their own password.
func setPassword(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		tx, err := db.BeginTxx(ctx, nil)
		if err != nil {
			return reportInt(err, w, r)
		}
		defer tx.Rollback()

		user, err := userByAuth(r, site.JWTSecret, tx)
		if err != nil {
			return report(http.StatusForbidden, false, "JWT authentication failure", err, false, w, r)
		}

		var req SetPasswordRequest
		dec := json.NewDecoder(r.Body)
		if err = dec.Decode(&req); err != nil {
			return report(http.StatusForbidden, false, "bad SetPasswrdRequest", err, false, w, r)
		}

		if req.Username != user.Name && !user.Admin {
			return report(http.StatusForbidden, false, "admin access required", nil, false, w, r)
		}

		username := req.Username
		if username == "" {
			username = user.Name
		}

		u, err := models.UserByName(username, tx)
		if err != nil {
			return reportInt(err, w, r)
		}

		if u == nil {
			return report(http.StatusNotFound, "user not found", "user not found", nil, false, w, r)
		}

		if !user.Admin || user.ID == u.ID {
			if !utils.ComparePassword(site.Config.PasswordSecret, req.OldPassword,
				u.Password) {
				return report(http.StatusForbidden, "User or password failure",
					"setPassword: old password mismatch",
					nil, false, w, r)
			}
		}

		newHashedPw, err := utils.HashPassword(site.Config.PasswordSecret, req.NewPassword)
		if err != nil {
			return reportInt(err, w, r)
		}
		u.Password = newHashedPw

		if err = u.Update(tx); err != nil {
			return reportInt(err, w, r)
		}

		if err = tx.Commit(); err != nil {
			return reportInt(err, w, r)
		}

		log.Printf("Password changed for %s by %s [%s]", u.Name, user.Name, r.Host)
		return respondJSON(true, w)
	}
}
