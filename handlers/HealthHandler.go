package handlers

import (
	"fmt"
	"net/http"
)

func HealthHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(200)

		body := "OK"

		fmt.Fprint(w, body)
	})
}
