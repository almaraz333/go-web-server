package handlers

import (
	"html/template"
	"net/http"
)

func MetricsHandler(hits *int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		data := struct {
			Hits int
		}{
			Hits: (*hits),
		}

		tmpl, err := template.ParseFiles("html/metrics.html")

		if err != nil {
			http.Error(w, "TEMPLATE ERROR", http.StatusInternalServerError)
		}

		execErr := tmpl.Execute(w, data)

		if execErr != nil {
			http.Error(w, "TEMPLATE EXECUTION ERROR", http.StatusInternalServerError)
		}
	})
}
