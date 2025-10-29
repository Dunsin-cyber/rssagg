package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/Dunsin-cyber/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)




func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
		
	}
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("invalid request payload: %v", err))
		return
	}
		feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID: uuid.New(),
		FeedID: params.FeedID,
		UserID: user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to create feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))
} 


func  (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	
		userID := user.ID
		feeds, err := apiCfg.DB.GetFeedFollows(r.Context(), userID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to get feed follow: %v", err))
		return
	}
	
		respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feeds))

		
}


func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	// get ID from param
	feedFollowIdStr := chi.URLParam(r, "feedFollowID")
	feedFollowIdParsed, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, 400, "invalid feed follow ID")
		return
	}

	 err = apiCfg.DB.DeleteFeedFollows(r.Context(), db.DeleteFeedFollowsParams{
		UserID: user.ID,
		ID: feedFollowIdParsed,
	})

	if  err != nil {
		respondWithError(w, 400, "could not delete feed follows")
		return
	}
	
	respondWithJSON(w, 200, struct{}{})
}