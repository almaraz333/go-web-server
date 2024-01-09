package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
)

func Chirp(id *int, db database.DBStructure, realDB database.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		chirp, _ := db.CreateChirp(cleanedString, *id)

		realDB.WriteDB(db)

		utils.RespondWithJSON(w, 201, chirp)

		*id++
	})
}
