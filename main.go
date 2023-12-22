package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/almaraz333/go-web-server/handlers"
	"github.com/go-chi/chi/v5"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type apiConfig struct {
	fileServerHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits++

		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricsReset() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits = 0

		w.WriteHeader(200)

		fmt.Fprint(w, "OK")
	})
}

func main() {
	apiConfig := apiConfig{}

	port := "8000"
	router := chi.NewRouter()
	api := chi.NewRouter()
	adminRouter := chi.NewRouter()

	router.Handle("/app/*", apiConfig.middlewareMetricsInc(handlers.AppHandler()))

	router.Handle("/app", apiConfig.middlewareMetricsInc(handlers.AppHandler()))

	router.Handle("/app/assets", apiConfig.middlewareMetricsInc(handlers.LogoHandler()))

	router.Handle("/assets/logo.png", apiConfig.middlewareMetricsInc(handlers.LogoImageHandler()))

	api.Get("/healthz", handlers.HealthHandler())

	api.Handle("/reset", apiConfig.metricsReset())

	api.Post("/validate_chirp", handlers.ValididateChirpHandler())

	adminRouter.Get("/metrics", handlers.MetricsHandler(&apiConfig.fileServerHits))

	corsMux := middlewareCors(router)
	router.Mount("/api", api)
	router.Mount("/admin", adminRouter)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	fmt.Printf("Starting Server on port %v ...\n", port)

	log.Fatal(srv.ListenAndServe())
}
