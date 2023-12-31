package handlers

import (
	"net/http"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
)

func GetChirps(db database.DBStructure) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirps, _ := db.GetChirps()

		utils.RespondWithJSON(w, 200, chirps)
	})
}
