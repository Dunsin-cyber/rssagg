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

type Feed struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Url       string `json:"url"`
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func databaseFeedToFeed(dbFeed db.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		UserID: dbFeed.UserID,
		Url: dbFeed.Url,
		Name: dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func databaseFeedsToFeeds(dbFeeds []db.Feed) []Feed {
	var feeds []Feed
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

func databaseFeedFollowToFeedFollow(dbFeedFollow db.FeedFollow) FeedFollow {
return FeedFollow{
		ID: dbFeedFollow.ID,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []db.FeedFollow) []FeedFollow {
		var feed_follows []FeedFollow
	for _, dbFeedFollow := range dbFeedFollows {
		feed_follows = append(feed_follows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feed_follows	
}