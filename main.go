package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/almaraz333/go-web-server/database"
	"github.com/almaraz333/go-web-server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
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
	secret         string
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
	godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")

	dbg := flag.Bool("debug", false, "Enable Debug Mode")
	flag.Parse()

	if *dbg {
		e := os.Remove("./DB.json")

		if e != nil {
			fmt.Println("")
		}
	}

	apiConfig := apiConfig{}

	apiConfig.secret = jwtSecret

	db, err := database.NewDB("./DB.json")

	dbStruct, loadDBError := db.LoadDB()

	if err != nil || loadDBError != nil {
		log.Fatalln("Could not create DB")
	}

	id := 1

	userId := 1

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

	api.Post("/chirps", handlers.Chirp(&id, dbStruct, *db))

	api.Get("/chirps", handlers.GetChirps(dbStruct))

	api.Get("/chirps/{chirpID}", handlers.GetChirpById(dbStruct))

	api.Post("/users", handlers.CreateUserHandler(&userId, dbStruct, *db))

	api.Post("/login", handlers.LoginHandler(dbStruct, *db, apiConfig.secret))

	api.Put("/users", handlers.UpdateUserHandler(dbStruct, *db, apiConfig.secret))

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
