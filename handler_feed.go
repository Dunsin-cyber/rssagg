package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/Dunsin-cyber/rssagg/internal/database"
	"github.com/google/uuid"
)




func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameter struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("invalid request payload: %v", err))
		return
	}
		feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID: uuid.New(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to create feed: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}


func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if  err != nil {
		respondWithError(w, 400, "could not get all users")
		return
	}
	
	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}