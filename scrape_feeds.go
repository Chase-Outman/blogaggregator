package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Chase-Outman/blogaggregator/internal/database"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feed.ID,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched:%w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Printf("Printing titles of %s:\n", feed.Name)
	fmt.Println("================================")
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Title:%s\n", item.Title)
	}
	return nil
}
