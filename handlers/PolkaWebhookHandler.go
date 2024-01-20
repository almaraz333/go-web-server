package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
)

func PolkaWebhookHandler(db *database.DBStructure) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		apiKey := strings.Split(authHeader, " ")

		if len(apiKey) < 2 {
			utils.RespondWithError(w, 401, "API key not provided")
			return
		}

		envAPIKey := os.Getenv("POLKA_KEY")

		if apiKey[1] != envAPIKey {
			utils.RespondWithError(w, 401, "API key not found")
			return
		}

		type BodyStruct struct {
			Event string `json:"event"`
			Data  struct {
				UserID int `json:"user_id"`
			} `json:"data"`
		}

		decoder := json.NewDecoder(r.Body)
		body := BodyStruct{}
		err := decoder.Decode(&body)

		if err != nil {
			utils.RespondWithError(w, 401, "Could not decode request body")
		}

		if body.Event != "user.upgraded" {
			utils.RespondWithJSON(w, 200, database.Chirp{})
			return
		}

		user, ok := db.Users[body.Data.UserID]

		if !ok {
			utils.RespondWithError(w, 404, "Cannot Find Error")
			return
		}

		user.IsChirpyRed = true
		db.Users[body.Data.UserID] = user

		utils.RespondWithJSON(w, 200, database.Chirp{})
	})
}
