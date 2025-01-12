package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
	"github.com/google/uuid"
)

func HandlerFollow(s *types.State, cmd types.Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("follow usage: gator follow <url>")
	}

	url := cmd.Arguments[0]
	feed, err := s.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed follow created successfully:")
	printFeedFollow(feedFollow)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeedFollow(feed database.CreateFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* User ID:          %s\n", feed.UserID)
	fmt.Printf("* Feed ID:           %s\n", feed.FeedID)
}
