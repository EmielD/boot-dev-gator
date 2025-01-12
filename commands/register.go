package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
	"github.com/google/uuid"
)

func HandlerRegister(s *types.State, cmd types.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New(("username expected for register command"))
	}

	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Arguments[0],
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	s.Config.SetUser(user.Name)
	fmt.Printf("User has been created: %v", user)

	return nil
}
