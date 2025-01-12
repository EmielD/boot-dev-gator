package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *types.State, cmd types.Command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: gator addfeed *name* *url*")
	}

	feedName := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]

	result, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    result.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully:")
	printFeed(result)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
