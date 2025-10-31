package main

import (
	"context"
	"log"
	"sync"
	"time"

	db "github.com/Dunsin-cyber/rssagg/internal/database"
)


func startScrapping(
	db *db.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {

	log.Printf("starting scrapper with concurrency %d and time between requests %s", concurrency, timeBetweenRequest.String())

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
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

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}


func scrapeFeed(db *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	defer wg.Done()

	 err := db.MarkFeedAsFetched(context.Background(), feed.ID)

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
		log.Println("Found post:", item.Title, "on feed:", feed.Name)
	}

	log.Printf("successfully scraped feed %s with %d items", feed.Name, len(rssFeed.Channel.Item))
}