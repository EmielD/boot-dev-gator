package commands

import (
	"context"
	"fmt"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
)

func HandlerFollowing(s *types.State, cmd types.Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}

	return nil
}
