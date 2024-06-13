package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PFrek/steam-affinity/internal/urls"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_ = os.Getenv("STEAM_APIKEY")
	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	addr := ":" + port
	server := http.Server{
		Handler: mux,
		Addr:    addr,

		ReadHeaderTimeout: 2 * time.Second,
	}

	apiV1baseURL := urls.URLChain{URL: "/api/v1/users"}
	mux.HandleFunc(apiV1baseURL.URL, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hi"))
	})

	log.Println("Starting server on port", port)
	server.ListenAndServe()
}
