package commands

import (
	"errors"
	"fmt"

	"github.com/emield/gator/internal/types"
)

func HandlerLogin(s *types.State, cmd types.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("username expected for login command")
	}

	_, err := s.Config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
