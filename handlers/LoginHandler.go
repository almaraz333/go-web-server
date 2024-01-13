package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(db database.DBStructure, realDB database.DB, secret string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		type loginStruct struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		decoder := json.NewDecoder(r.Body)
		body := loginStruct{}
		err := decoder.Decode(&body)

		if err != nil {
			utils.RespondWithError(w, 500, "Something went wrong")
			return
		}

		existingUser := database.User{}

		// PROBABLY NOT GOOD DO NOT DO IN PROD - JUST TESTING
		for _, val := range db.Users {
			if val.Email == body.Email {
				existingUser = val
				break
			}
		}

		compareErr := bcrypt.CompareHashAndPassword(existingUser.Password, []byte(body.Password))

		if compareErr != nil {
			utils.RespondWithError(w, 401, "Wrong password")
			return
		}

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			Issuer:    "chiry-access",
			Subject:   strconv.Itoa(existingUser.Id),
		})

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(24*60))),
			Issuer:    "chirpy-refresh",
			Subject:   strconv.Itoa(existingUser.Id),
		})

		accessJwt, accessSignErr := accessToken.SignedString([]byte(secret))

		refreshJwt, signErr := refreshToken.SignedString([]byte(secret))

		if signErr != nil || accessSignErr != nil {
			utils.RespondWithError(w, 401, "Could not sign token")
			return
		}

		userRes := struct {
			Id           int    `json:"id"`
			Email        string `json:"email"`
			Token        string `json:"token"`
			Refreshtoken string `json:"refresh_token"`
		}{
			Id:           existingUser.Id,
			Email:        existingUser.Email,
			Token:        accessJwt,
			Refreshtoken: refreshJwt,
		}

		utils.RespondWithJSON(w, 200, userRes)
	})
}
