package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	db "github.com/Dunsin-cyber/rssagg/internal/database"
	"github.com/google/uuid"
)


func startScrapping(
	DB *db.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {

	log.Printf("starting scrapper with concurrency %d and time between requests %s", concurrency, timeBetweenRequest.String())

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := DB.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			log.Printf("failed to get feeds to fetch: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(DB, wg, feed)
		}
		wg.Wait()
	}
}


func scrapeFeed(DB *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	defer wg.Done()

	 err := DB.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("failed to mark feed %s as fetched: %v", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	
	if err != nil {
		log.Printf("failed to parse feed %s: %v", feed.ID, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("couldn't parse date %v with err %v", item.PubDate, err)
			continue 
		}
		_, err = DB.CreatePost(context.Background(), db.CreatePostParams{
			ID:uuid.New(),
			CreatedAt:time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid: item.Description != "",
			},
			Url: item.Link,
			PublishedAt: pubAt,
			FeedID: feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Failed to create post %v with err %v", item.Title, err )
			continue
		}

	}

	log.Printf("successfully scraped feed %s with %d items", feed.Name, len(rssFeed.Channel.Item))
}