package main

import (
	"net/http"

	"github.com/Dunsin-cyber/rssagg/internal/auth"
	db "github.com/Dunsin-cyber/rssagg/internal/database"
)


type authedHandler func( http.ResponseWriter,  *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth( handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, "missing or invalid api key")
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, "could not get user by api key")
		return
	}
	handler(w, r, user)
}
}