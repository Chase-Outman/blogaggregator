package main

import (
	"context"
	"fmt"

	"github.com/Chase-Outman/blogaggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	feed_names, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feeds for user: %w", err)
	}

	for _, name := range feed_names {
		fmt.Println(name.FeedName)
	}

	return nil
}
