package api

import (
	"cmp"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"slices"
	"strings"
	"sync"
)

type ApiConfig struct {
	SteamApiKey string

	FriendsListCache Cache[FriendsList]
	SummariesCache   Cache[Summaries]
	OwnedGamesCache  Cache[OwnedGames]
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	responseStruct := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}

	data, err := json.Marshal(responseStruct)
	if err != nil {
		log.Println("Failed to marshal error body:", err)
		data = []byte("[invalid err body]")
	}

	w.Write(data)
}

func RespondWithJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data, err := json.Marshal(body)
	if err != nil {
		log.Println("Failed to marshal response body:", err)
		data = []byte(`"Invalid Body"`)
	}

	w.Write(data)
}

func logRequest(req *http.Request, statusCode int) {
	log.Printf("%s %s - %d", req.Method, req.URL, statusCode)
}

func (config *ApiConfig) GetFriendsHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("steamid")
	if id == "" {
		logRequest(req, 400)
		RespondWithError(w, 400, "Query param 'steamid' is required")
		return
	}

	friends, err := config.GetFriendList(id)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			logRequest(req, 400)
			RespondWithError(w, 400, err.Error())
			return
		}

		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	friendsIDs := friends.ExtractIds()

	summaries, err := config.GetPlayerSummaries(friendsIDs)
	if err != nil {
		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	logRequest(req, 200)
	RespondWithJSON(w, 200, summaries)
}

func (config *ApiConfig) GetSummariesHandler(w http.ResponseWriter, req *http.Request) {
	ids := req.URL.Query().Get("steamids")
	if ids == "" {
		logRequest(req, 400)
		RespondWithError(w, 400, "Query param 'steamids' is required")
		return
	}

	summaries, err := config.GetPlayerSummaries(strings.Split(ids, ","))
	if err != nil {
		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	logRequest(req, 200)
	RespondWithJSON(w, 200, summaries)
}

func (config *ApiConfig) GetOwnedGamesHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("steamid")
	if id == "" {
		logRequest(req, 400)
		RespondWithError(w, 400, "Query param 'steamid' is required")
		return
	}

	ownedGames, err := config.GetOwnedGames(id)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			logRequest(req, 400)
			RespondWithError(w, 400, err.Error())
			return
		}

		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	logRequest(req, 200)
	RespondWithJSON(w, 200, ownedGames)
}

func (config *ApiConfig) GetComparisonHandler(w http.ResponseWriter, req *http.Request) {
	player1ID := req.URL.Query().Get("player1")
	player2ID := req.URL.Query().Get("player2")
	if player1ID == "" {
		logRequest(req, 400)
		RespondWithError(w, 400, "Query param 'player1' is required")
		return
	}
	if player2ID == "" {
		logRequest(req, 400)
		RespondWithError(w, 400, "Query param 'player2' is required")
		return
	}

	listGames := false
	listGamesQuery := req.URL.Query().Get("listGames")
	if listGamesQuery == "true" {
		listGames = true
	}

	player1Games, err := config.GetOwnedGames(player1ID)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			logRequest(req, 400)
			RespondWithError(w, 400, err.Error())
			return
		}

		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	player2Games, err := config.GetOwnedGames(player2ID)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			logRequest(req, 400)
			RespondWithError(w, 400, err.Error())
			return
		}

		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	result := player1Games.CompareOwnedGames(player2Games, listGames)

	logRequest(req, 200)
	RespondWithJSON(w, 200, result)
}

func (config *ApiConfig) GetAffinityRanking(w http.ResponseWriter, req *http.Request) {
	steamid := req.URL.Query().Get("steamid")
	if steamid == "" {
		logRequest(req, 400)
		RespondWithError(w, 400, "Query param 'steamid' is required")
		return
	}

	listGames := false
	listGamesQuery := req.URL.Query().Get("listGames")
	if listGamesQuery == "true" {
		listGames = true
	}

	ownedGames, err := config.GetOwnedGames(steamid)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			logRequest(req, 400)
			RespondWithError(w, 400, err.Error())
			return
		}

		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	friendsList, err := config.GetFriendList(steamid)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			logRequest(req, 400)
			RespondWithError(w, 400, err.Error())
			return
		}

		logRequest(req, 500)
		RespondWithError(w, 500, err.Error())
		return
	}

	wg := sync.WaitGroup{}
	ch := make(chan CompareResult, len(friendsList.Friends))

	for _, friend := range friendsList.Friends {
		wg.Add(1)

		go func(friend Friend, ch chan CompareResult) {
			defer wg.Done()

			friendsGames, err := config.GetOwnedGames(friend.SteamID)
			if err != nil {
				if errors.Is(err, InvalidSteamIDError{}) {
					logRequest(req, 400)
					RespondWithError(w, 400, err.Error())
					return
				}

				logRequest(req, 500)
				RespondWithError(w, 500, err.Error())
				return
			}

			result := ownedGames.CompareOwnedGames(friendsGames, listGames)

			ch <- result
		}(friend, ch)
	}

	wg.Wait()

	results := []CompareResult{}
	for range len(friendsList.Friends) {
		result := <-ch
		results = append(results, result)
	}

	slices.SortFunc(results, func(a CompareResult, b CompareResult) int {
		return cmp.Compare(b.Affinity, a.Affinity)
	})

	response := struct {
		Ranking []CompareResult `json:"ranking"`
	}{
		Ranking: results,
	}

	logRequest(req, 200)
	RespondWithJSON(w, 200, response)
}
