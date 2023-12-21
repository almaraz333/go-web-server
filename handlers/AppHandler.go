package handlers

import (
	"net/http"
)

func AppHandler() http.Handler {
	fileServer := http.FileServer(http.Dir("html/homepage"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache")
		http.StripPrefix("/app", fileServer).ServeHTTP(w, r)
	})
}
