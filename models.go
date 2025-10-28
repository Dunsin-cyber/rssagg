package main

import (
	"time"

	db "github.com/Dunsin-cyber/rssagg/internal/database"
	"github.com/google/uuid"
)


type User struct {
	ID			uuid.UUID `json:"id"`
	Name		string `json:"name"`
	CreatedAt	time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
	ApiKey		string `json:"api_key"`
}

func databaseToUser(dbUser db.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey: dbUser.ApiKey,
	}
}