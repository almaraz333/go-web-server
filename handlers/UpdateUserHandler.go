package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
	"github.com/golang-jwt/jwt/v5"
)

func UpdateUserHandler(db database.DBStructure, realDB database.DB, secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		splitAuth := strings.Split(authHeader, " ")

		type BodyStruct struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		decoder := json.NewDecoder(r.Body)
		body := BodyStruct{}
		err := decoder.Decode(&body)

		if err != nil {
			utils.RespondWithError(w, 400, "Could not decode request body")
		}

		parsedJwt, jwtParseError := jwt.ParseWithClaims(splitAuth[1], &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })

		if jwtParseError != nil {
			utils.RespondWithError(w, 401, "Unauthorized")
			return
		}

		sub, subErr := parsedJwt.Claims.GetSubject()

		if subErr != nil {
			utils.RespondWithError(w, 400, "Invalid user ID")
		}

		intSub, _ := strconv.Atoi(sub)

		fmt.Println(db.Users)

		updatedUser, updateUserErr := realDB.UpdateUser(intSub, body.Email, db, body.Password)

		updatedUserStruct := struct {
			Email string `json:"email"`
			Id    int    `json:"id"`
		}{
			Email: updatedUser.Email,
			Id:    updatedUser.Id,
		}

		fmt.Println(db.Users)

		if updateUserErr != nil {
			utils.RespondWithError(w, 500, "Could not update user")
			return
		}

		utils.RespondWithJSON(w, 200, updatedUserStruct)
	}
}
