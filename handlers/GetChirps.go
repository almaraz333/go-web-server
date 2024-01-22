package handlers

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/utils"
)

func GetChirps(db *database.DBStructure) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorId := r.URL.Query().Get("author_id")
		sortOrder := r.URL.Query().Get("sort")

		chirps := db.Chirps[:]

		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].Id < chirps[j].Id
		})

		if sortOrder == "desc" {
			sort.Slice(chirps, func(i, j int) bool {
				return chirps[i].Id > chirps[j].Id
			})
		}

		if authorId != "" {
			chirpsByAuthor := make([]database.Chirp, 0)

			intAuthorId, _ := strconv.Atoi(authorId)

			for _, chirp := range db.Chirps {
				if chirp.AuthorId == intAuthorId {
					chirpsByAuthor = append(chirpsByAuthor, chirp)
				}

			}

			utils.RespondWithJSON(w, 200, chirpsByAuthor)
			return
		}

		utils.RespondWithJSON(w, 200, db.Chirps)
	})
}
