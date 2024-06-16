package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/PFrek/steam-affinity/internal/api"
	"github.com/joho/godotenv"
)

func TryAtoi(s string, defaultValue int) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return result
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	steamApiKey := os.Getenv("STEAM_APIKEY")
	port := os.Getenv("PORT")

	friendsCacheRenew := time.Duration(TryAtoi(os.Getenv("FRIENDS_CACHE_RENEW"), 5))
	playersCacheRenew := time.Duration(TryAtoi(os.Getenv("SUMMARIES_CACHE_RENEW"), 1440))
	gamesCacheRenew := time.Duration(TryAtoi(os.Getenv("GAMES_CACHE_RENEW"), 30))

	config := api.ApiConfig{
		SteamApiKey: steamApiKey,
		FriendsListCache: api.Cache[api.FriendsList]{
			Cache:      map[string]api.CacheEntry[api.FriendsList]{},
			CacheRenew: friendsCacheRenew * time.Minute,
		},
		PlayersCache: api.Cache[api.Player]{
			Cache:      map[string]api.CacheEntry[api.Player]{},
			CacheRenew: playersCacheRenew * time.Minute,
		},
		OwnedGamesCache: api.Cache[api.OwnedGames]{
			Cache:      map[string]api.CacheEntry[api.OwnedGames]{},
			CacheRenew: gamesCacheRenew * time.Minute,
		},
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
