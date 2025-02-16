package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Chase-Outman/blogaggregator/internal/database"
	"github.com/google/uuid"
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
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {

		}

	}
	return nil
}
