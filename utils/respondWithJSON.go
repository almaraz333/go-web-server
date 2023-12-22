package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(payload)

	if err != nil {
		RespondWithError(w, 500, "Something went wrong")
		return
	}

	w.Write(data)
}
