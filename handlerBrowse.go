package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Chase-Outman/blogaggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) == 1 {
		l, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("imporper aguments: %w", err)
		}
		limit = l
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to get posts for current user: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Link to post: %s\n", post.Url)
	}
	return nil

}
