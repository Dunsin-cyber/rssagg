package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	db "github.com/Dunsin-cyber/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameter{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("invalid request payload: %v", err))
		return
	}
		user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("failed to create user: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user db.User) {
	respondWithJSON(w, 200, databaseToUser(user))

}


func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user db.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), db.GetPostsForUserParams{
		UserID:user.ID,
		Limit:10,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
	}
	respondWithJSON(w, 200, databasePostsToPosts(posts))

}

func (apiCfg *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())
	if  err != nil {
		respondWithError(w, 400, "could not get all users")
		return
	}
	var resp []User
	for _, u := range users {
		resp = append(resp, databaseToUser(u))
	}
	respondWithJSON(w, 200, resp)
}