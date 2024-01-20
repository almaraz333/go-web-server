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

func Chirp(secret string, id *int, db *database.DBStructure, realDB database.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authTokenFromHeader := strings.Split(authHeader, " ")[1]

		parsedJwt, jwtParseError := jwt.ParseWithClaims(authTokenFromHeader, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })

		if jwtParseError != nil {
			utils.RespondWithError(w, 401, "Unauthorized")
			return
		}

		userId, userIdParseError := parsedJwt.Claims.GetSubject()

		if userIdParseError != nil {
			utils.RespondWithError(w, 401, "Unauthorized")
			return
		}

		intSub, _ := strconv.Atoi(userId)

		type bodyStruct struct {
			Body string `json:"body"`
		}

		decoder := json.NewDecoder(r.Body)
		body := bodyStruct{}
		err := decoder.Decode(&body)

		if err != nil {
			utils.RespondWithError(w, 500, "Something went wrong")
			return
		}

		if len(body.Body) > 140 {
			utils.RespondWithError(w, 400, "Chirp is too long")
			return
		}

		cleanedString := utils.CleanChirp(body.Body)

		chirp, _ := db.CreateChirp(cleanedString, intSub, *id)

		realDB.WriteDB(*db)

		fmt.Println(db.Chirps)

		*id++

		utils.RespondWithJSON(w, 201, chirp)
	})
}
