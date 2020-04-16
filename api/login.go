package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/terotoi/marque/core"
	"github.com/terotoi/marque/jwt"
	"github.com/terotoi/marque/models"
	"github.com/terotoi/marque/utils"
	"github.com/jmoiron/sqlx"
)

func login(site *core.Site, db *sqlx.DB) AnonAPIfunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		dec := json.NewDecoder(r.Body)

		var loginReq struct {
			Username string
			Password string
		}

		err := dec.Decode(&loginReq)
		if err != nil {
			return err
		}

		user, err := models.UserByName(loginReq.Username, db)
		if err != nil {
			return err
		}

		if user == nil {
			return fmt.Errorf("user not found")
		}

		hashedPw := utils.MakeHash(site.Config.PasswordSecret, loginReq.Password)
		if user.Password == nil || *user.Password != hashedPw {
			return fmt.Errorf("password mismatch")
		}

		token, err := jwt.CreateToken(site.JWTSecret,
			jwtAuth{
				strconv.Itoa(int(user.ID)),
			})

		if err != nil {
			return err
		}

		response := struct {
			Token string
		}{
			token,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return nil
	}
}
