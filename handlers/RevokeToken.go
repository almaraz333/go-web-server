package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
)

func RevokeTokenHandler(db database.DBStructure) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authTokenFromHeader := strings.Split(authHeader, " ")[1]

		db.RevokedTokens[authTokenFromHeader] = time.Now()

		utils.RespondWithJSON(w, 200, struct{}{})
	}
}
