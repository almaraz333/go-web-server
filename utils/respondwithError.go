package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorStruct struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	log.Printf("Error with code %v: %v", code, msg)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	errorBody := errorStruct{
		Error: msg,
	}

	data, _ := json.Marshal(errorBody)
	w.Write(data)
}
