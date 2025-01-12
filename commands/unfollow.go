package commands

import (
	"context"
	"errors"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
)

func HandleUnfollow(s *types.State, cmd types.Command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return errors.New("unfollow usage: unfollow <url>")
	}

	err := s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Name: user.Name,
		Url:  cmd.Arguments[0],
	})
	if err != nil {
		return err
	}

	return nil
}
