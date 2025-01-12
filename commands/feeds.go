package commands

import (
	"context"
	"fmt"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
)

func HandlerFeeds(s *types.State, cmd types.Command) error {

	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, f := range feeds {
		user, err := s.Db.GetUserById(context.Background(), f.UserID)
		if err != nil {
			return err
		}
		printFeedWithUsername(f, user.Name)
	}

	return nil
}

func printFeedWithUsername(feed database.Feed, userName string) {
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* Username:        %s\n", userName)
}
