package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const steamBaseURL = "http://api.steampowered.com/ISteamUser/"

type Friend struct {
	SteamID      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

type FriendsList struct {
	Friends []Friend `json:"friends"`
}

type FriendsListResponse struct {
	FriendsList FriendsList `json:"friendslist"`
}

func (config *ApiConfig) GetFriendList(steamid string) (FriendsList, error) {
	base, err := url.Parse(steamBaseURL)
	if err != nil {
		return FriendsList{}, errors.New("Failed to parse steam API URL")
	}
	base = base.JoinPath("GetFriendList", "v0001/")

	query := url.Values{}
	query.Set("key", config.SteamApiKey)
	query.Set("steamid", steamid)
	query.Set("relationship", "friend")

	base.RawQuery = query.Encode()

	log.Println("Sending request to steam API at:", base.String())

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

type InvalidSteamIDError struct{}

func (err InvalidSteamIDError) Error() string {
	return "Invalid steamid"
}
