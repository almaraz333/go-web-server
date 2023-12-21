package handlers

import (
	"net/http"
)

func LogoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/logo.html")
	})
}
