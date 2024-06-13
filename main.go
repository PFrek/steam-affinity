package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PFrek/steam-affinity/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	steamApiKey := os.Getenv("STEAM_APIKEY")
	port := os.Getenv("PORT")

	config := api.ApiConfig{
		SteamApiKey: steamApiKey,
	}

	mux := http.NewServeMux()

	addr := ":" + port
	server := http.Server{
		Handler: mux,
		Addr:    addr,

		ReadHeaderTimeout: 2 * time.Second,
	}

	apiV1baseURL := "/api/v1"

	mux.HandleFunc(apiV1baseURL+"/friends", config.GetFriendsHandler)
	mux.HandleFunc(apiV1baseURL+"/summaries", config.GetSummariesHandler)
	mux.HandleFunc(apiV1baseURL+"/ownedGames", config.GetOwnedGamesHandler)
	mux.HandleFunc(apiV1baseURL+"/ownedGames/compare", config.GetComparisonHandler)
	mux.HandleFunc(apiV1baseURL+"/friends/ranking", config.GetAffinityRanking)

	log.Println("Starting server on port", port)
	server.ListenAndServe()
}
