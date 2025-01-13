package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/emield/gator/internal/types"
)

func HandlerLogin(s *types.State, cmd types.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("username expected for login command")
	}

	user, err := s.Db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("user does not exist, register the user first using: register <name>")
	}

	_, err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
