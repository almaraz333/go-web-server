package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
)

func CreateUserHandler(id *int, db database.DBStructure, realDB database.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type bodyStruct struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		decoder := json.NewDecoder(r.Body)
		body := bodyStruct{}
		err := decoder.Decode(&body)

		if err != nil {
			utils.RespondWithError(w, 500, "Something went wrong")
			return
		}

		user, createUserError := db.CreateUser(body.Email, *id, body.Password)

		if createUserError != nil {
			utils.RespondWithError(w, 400, createUserError.Error())
			return
		}

		realDB.WriteDB(db)

		*id++

		utils.RespondWithJSON(w, 201, user)

	})
}
