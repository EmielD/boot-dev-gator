package commands

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
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
