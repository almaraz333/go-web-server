package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
	"github.com/golang-jwt/jwt/v5"
)

func RefreshHandler(secret string, database database.DBStructure) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authTokenFromHeader := strings.Split(authHeader, " ")[1]

		parsedJwt, jwtParseError := jwt.ParseWithClaims(authTokenFromHeader, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })

		if jwtParseError != nil {
			utils.RespondWithError(w, 401, "Unauthorized")
			return
		}

		issuer, issuerParseError := parsedJwt.Claims.GetIssuer()
		userId, userIdParseError := parsedJwt.Claims.GetSubject()

		if userIdParseError != nil {
			utils.RespondWithError(w, 401, "Unauthorized")
			return
		}

		intSub, _ := strconv.Atoi(userId)

		if issuer != "chirpy-refresh" || issuerParseError != nil {
			utils.RespondWithError(w, 401, "Not a refresh token")
			return
		}

		_, found := database.RevokedTokens[authTokenFromHeader]

		if found {
			utils.RespondWithError(w, 401, "Cannot used Revoked token")
			return
		}

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			Issuer:    "chiry-access",
			Subject:   strconv.Itoa(intSub),
		})

		jwt, signErr := accessToken.SignedString([]byte(secret))

		if signErr != nil {
			utils.RespondWithError(w, 400, "Could not sign token")
			return
		}

		resStruct := struct {
			Token string `json:"token"`
		}{
			Token: jwt,
		}

		utils.RespondWithJSON(w, 200, resStruct)
	}
}
