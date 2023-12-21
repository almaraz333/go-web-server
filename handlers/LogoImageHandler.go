package handlers

import (
	"net/http"
)

func LogoImageHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/logo.png")
	})
}
