package commands

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("username expected for login command")
	}

	user, err := s.Db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	_, err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
