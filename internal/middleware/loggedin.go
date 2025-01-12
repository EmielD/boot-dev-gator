package middleware

import (
	"context"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
)

func LoggedIn(handler func(s *types.State, cmd types.Command, user database.User) error) func(*types.State, types.Command) error {
	return func(s *types.State, c types.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
