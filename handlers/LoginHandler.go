package handlers

import (
	"encoding/json"
	"fmt"
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
			Email            string `json:"email"`
			Password         string `json:"password"`
			ExpiresInSeconds int    `json:"expires_in_seconds"`
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

		fmt.Printf("existingUser: %v\n", existingUser)

		compareErr := bcrypt.CompareHashAndPassword(existingUser.Password, []byte(body.Password))

		if compareErr != nil {
			utils.RespondWithError(w, 401, "Wrong password")
			return
		}

		expTime := 50000

		if body.ExpiresInSeconds > 0 {
			expTime = body.ExpiresInSeconds
		}

		tokenPreSigned := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			Issuer:    "chirpy",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expTime))),
			Subject:   strconv.Itoa(existingUser.Id),
		})

		jwt, signErr := tokenPreSigned.SignedString([]byte(secret))

		if signErr != nil {
			utils.RespondWithError(w, 400, "Could not sign token")
			return
		}

		userRes := struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
			Token string `json:"token"`
		}{
			Id:    existingUser.Id,
			Email: existingUser.Email,
			Token: jwt,
		}

		utils.RespondWithJSON(w, 200, userRes)
	})
}
