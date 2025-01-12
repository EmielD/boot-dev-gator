package commands

import (
	"context"
	"fmt"

	"github.com/emield/gator/internal/types"
)

func HandlerUsers(s *types.State, cmd types.Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		isCurrentString := ""
		if s.Config.Current_user_name == user.Name {
			isCurrentString = (" (current)\n")
		} else {
			isCurrentString = "\n"
		}
		fmt.Printf("* %s%s", user.Name, isCurrentString)
	}

	return nil
}
