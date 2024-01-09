package handlers

import (
	"net/http"
	"strconv"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
	"github.com/go-chi/chi/v5"
)

func GetChirpById(db database.DBStructure) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "chirpID")

		convertedId, conversionErr := strconv.Atoi(id)

		if conversionErr != nil {
			utils.RespondWithJSON(w, 404, database.DBStructure{})
		}

		chirp, err := db.GetChirpByID(convertedId)

		if err != nil {
			utils.RespondWithJSON(w, 404, database.DBStructure{})
		}

		utils.RespondWithJSON(w, 200, chirp)
	})
}
