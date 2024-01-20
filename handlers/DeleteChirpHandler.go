package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

func DeleteChirpHandler(db *database.DBStructure, secret string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "chirpID")

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

		convertedId, conversionErr := strconv.Atoi(id)

		if conversionErr != nil {
			utils.RespondWithJSON(w, 404, database.DBStructure{})
			return
		}

		for i, chirp := range db.Chirps {
			if chirp.Id == convertedId && chirp.AuthorId == intSub {
				db.Chirps = append(db.Chirps[:i], db.Chirps[i+1:]...)
				utils.RespondWithJSON(w, 200, database.Chirp{})
				return
			}
		}

		utils.RespondWithJSON(w, 403, database.Chirp{})
	})
}
