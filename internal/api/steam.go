package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

const steamBaseURL = "http://api.steampowered.com/"
const steamUserAPIURL = "ISteamUser/"
const steamPlayerAPIURL = "IPlayerService/"

type Friend struct {
	SteamID      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

type FriendsList struct {
	Friends []Friend `json:"friends"`
}

func (fl FriendsList) ExtractIds() []string {
	ids := []string{}

	for _, friend := range fl.Friends {
		ids = append(ids, friend.SteamID)
	}

	return ids
}

type FriendsListResponse struct {
	FriendsList FriendsList `json:"friendslist"`
}

func (config *ApiConfig) GetFriendList(steamid string) (FriendsList, error) {
	base, err := url.Parse(steamBaseURL)
	if err != nil {
		return FriendsList{}, errors.New("Failed to parse steam API URL")
	}
	base = base.JoinPath(steamUserAPIURL, "GetFriendList", "v0001/")

	query := url.Values{}
	query.Set("key", config.SteamApiKey)
	query.Set("steamid", steamid)
	query.Set("relationship", "friend")

	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return FriendsList{}, err
	}
	defer resp.Body.Close()

	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		html, err := io.ReadAll(resp.Body)
		if err == nil {
			log.Printf("%s\n", html)
		}
		return FriendsList{}, InvalidSteamIDError{}
	}

	decoder := json.NewDecoder(resp.Body)

	body := FriendsListResponse{}
	err = decoder.Decode(&body)
	if err != nil {
		return FriendsList{}, err
	}

	return body.FriendsList, nil
}

type Player struct {
	SteamID                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	PersonaName              string `json:"personaname"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
}

type Summaries struct {
	Players []Player `json:"players"`
}

type SummariesResponse struct {
	Response Summaries `json:"response"`
}

func (config *ApiConfig) GetPlayerSummaries(steamids []string) (Summaries, error) {
	base, err := url.Parse(steamBaseURL)
	if err != nil {
		return Summaries{}, errors.New("Failed to parse steam API URL")
	}
	base = base.JoinPath(steamUserAPIURL, "GetPlayerSummaries", "v0002/")

	query := url.Values{}
	query.Set("key", config.SteamApiKey)
	query.Set("steamids", strings.Join(steamids, ","))

	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return Summaries{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	body := SummariesResponse{}
	err = decoder.Decode(&body)
	if err != nil {
		return Summaries{}, err
	}

	return body.Response, nil
}

type Game struct {
	AppID      int    `json:"appid"`
	Name       string `json:"name"`
	ImgIconURL string `json:"img_icon_url"`
}

type OwnedGames struct {
	SteamID   string
	GameCount int    `json:"game_count"`
	Games     []Game `json:"games"`
}

type OwnedGamesResponse struct {
	Response OwnedGames `json:"response"`
}

func (config *ApiConfig) GetOwnedGames(steamid string) (OwnedGames, error) {
	base, err := url.Parse(steamBaseURL)
	if err != nil {
		return OwnedGames{}, errors.New("Failed to parse steam API URL")
	}
	base = base.JoinPath(steamPlayerAPIURL, "GetOwnedGames", "v0001/")

	query := url.Values{}
	query.Set("key", config.SteamApiKey)
	query.Set("steamid", steamid)
	query.Set("include_appinfo", "true")

	base.RawQuery = query.Encode()

	resp, err := http.Get(base.String())
	if err != nil {
		return OwnedGames{}, err
	}
	defer resp.Body.Close()

	if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		html, err := io.ReadAll(resp.Body)
		if err == nil {
			log.Printf("%s\n", html)
		}
		return OwnedGames{}, InvalidSteamIDError{}
	}

	decoder := json.NewDecoder(resp.Body)

	body := OwnedGamesResponse{}
	err = decoder.Decode(&body)
	if err != nil {
		return OwnedGames{}, err
	}

	body.Response.SteamID = steamid

	return body.Response, nil

}

type CompareResult struct {
	AffinityAvg   float64 `json:"affinity_avg"`
	Affinity1     float64 `json:"affinity_1"`
	Affinity2     float64 `json:"affinity_2"`
	Player1ID     string  `json:"player1ID"`
	Player2ID     string  `json:"player2ID"`
	Matches       int     `json:"matches"`
	MatchingGames []Game  `json:"matching_games"`
}

func (player1Games OwnedGames) CompareOwnedGames(player2Games OwnedGames, listGames bool) CompareResult {
	result := CompareResult{
		Player1ID: player1Games.SteamID,
		Player2ID: player2Games.SteamID,
	}

	result.MatchingGames = []Game{}
	for _, game := range player1Games.Games {
		if slices.Contains(player2Games.Games, game) {
			result.MatchingGames = append(result.MatchingGames, game)
		}
	}

	result.Affinity1 = GetAffinity(result.MatchingGames, player1Games)
	result.Affinity2 = GetAffinity(result.MatchingGames, player2Games)
	result.AffinityAvg = (result.Affinity1 + result.Affinity2) / 2.0
	diffFactor := (1.0 - math.Abs(result.Affinity1-result.Affinity2))
	result.AffinityAvg *= diffFactor

	result.Matches = len(result.MatchingGames)

	if !listGames {
		result.MatchingGames = nil
	}
	return result
}

func GetAffinity(matchingGames []Game, playerGames OwnedGames) float64 {
	return float64(len(matchingGames)) / float64(playerGames.GameCount)
}

type InvalidSteamIDError struct{}

func (err InvalidSteamIDError) Error() string {
	return "Invalid steamid"
}
