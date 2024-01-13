package utils

import (
	"encoding/json"
	"net/http"
)

type errorStruct struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	errorBody := errorStruct{
		Error: msg,
	}

	data, _ := json.Marshal(errorBody)
	w.Write(data)
}
