package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ApiConfig struct {
	SteamApiKey string
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

func (config *ApiConfig) GetFriendsHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("steamid")
	if id == "" {
		RespondWithError(w, 400, "Query param 'steamid' is required")
		return
	}

	friends, err := config.GetFriendList(id)
	if err != nil {
		if errors.Is(err, InvalidSteamIDError{}) {
			RespondWithError(w, 400, err.Error())
			return
		}

		RespondWithError(w, 500, err.Error())
		return
	}

	RespondWithJSON(w, 200, friends)
}
