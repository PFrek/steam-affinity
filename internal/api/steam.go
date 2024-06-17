package api

import (
	"cmp"
	"encoding/json"
	"errors"
	"io"
	"log"
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
	if config.FriendsListCache.IsCacheHit(steamid) {
		log.Println("FriendsList cache hit for", steamid)
		return config.FriendsListCache.Cache[steamid].Data, nil
	}

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

	config.FriendsListCache.UpdateCache(steamid, body.FriendsList)

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
	uncachedIDs := []string{}
	cachedPlayers := []Player{}
	for _, id := range steamids {
		if config.PlayersCache.IsCacheHit(id) {
			log.Println("Cache hit for id:", id)
			cachedPlayers = append(cachedPlayers, config.PlayersCache.ReadCache(id))
		} else {
			uncachedIDs = append(uncachedIDs, id)
		}
	}

	if len(uncachedIDs) == 0 {
		slices.SortFunc(cachedPlayers, func(a Player, b Player) int {
			return cmp.Compare(a.SteamID, b.SteamID)
		})
		return Summaries{
			Players: cachedPlayers,
		}, nil
	}

	joinedIDs := strings.Join(uncachedIDs, ",")

	base, err := url.Parse(steamBaseURL)
	if err != nil {
		return Summaries{}, errors.New("Failed to parse steam API URL")
	}
	base = base.JoinPath(steamUserAPIURL, "GetPlayerSummaries", "v0002/")

	query := url.Values{}
	query.Set("key", config.SteamApiKey)
	query.Set("steamids", joinedIDs)

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

	returnedPlayers := body.Response.Players
	for _, player := range returnedPlayers {
		log.Println("Adding player to cache with id:", player.SteamID)
		config.PlayersCache.UpdateCache(player.SteamID, player)
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
	if config.OwnedGamesCache.IsCacheHit(steamid) {
		log.Println("OwnedGames cache hit for", steamid)
		return config.OwnedGamesCache.Cache[steamid].Data, nil
	}

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

	config.OwnedGamesCache.UpdateCache(steamid, body.Response)

	return body.Response, nil

}

type CompareResult struct {
	Affinity          float64 `json:"affinity"`
	Similarity        float64 `json:"similarity"`
	Weight            float64 `json:"weight"`
	Player1ID         string  `json:"player1ID"`
	Player1Ratio      float64 `json:"player1Ratio"`
	Player2ID         string  `json:"player2ID"`
	Player2Ratio      float64 `json:"player2Ratio"`
	Player2GamesCount int     `json:"player2GamesCount"`
	Matches           int     `json:"matches"`
	MatchingGames     []Game  `json:"matching_games"`
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
	result.Matches = len(result.MatchingGames)

	if player1Games.GameCount > 0 {
		result.Player1Ratio = float64(result.Matches) / float64(player1Games.GameCount)
	}

	result.Player2GamesCount = player2Games.GameCount
	if player2Games.GameCount > 0 {
		result.Player2Ratio = float64(result.Matches) / float64(player2Games.GameCount)
	}

	if player1Games.GameCount+player2Games.GameCount > 0 {
		result.Similarity = float64(len(result.MatchingGames)) / float64(player1Games.GameCount+player2Games.GameCount)
		result.Weight = (2.0 * float64(player1Games.GameCount) * float64(player2Games.GameCount)) / (float64(player1Games.GameCount + player2Games.GameCount))
		result.Affinity = result.Similarity * result.Weight
	}

	if !listGames {
		result.MatchingGames = nil
	}

	return result
}

func GetAffinity(matchingGames []Game, playerGames OwnedGames) float64 {
	if playerGames.GameCount == 0 {
		return 0
	}
	return float64(len(matchingGames)) / float64(playerGames.GameCount)
}

type InvalidSteamIDError struct{}

func (err InvalidSteamIDError) Error() string {
	return "Invalid steamid"
}
