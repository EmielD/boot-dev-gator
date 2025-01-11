package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/emield/gator/internal/config"
	"github.com/emield/gator/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("could not run command, state is nil")
	}

	var command = c.commandMap[cmd.name]
	if command == nil {
		return errors.New("command does not exist")
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		isCurrentString := ""
		if s.config.Current_user_name == user.Name {
			isCurrentString = (" (current)\n")
		} else {
			isCurrentString = "\n"
		}
		fmt.Printf("* %s%s", user.Name, isCurrentString)
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	s.db.Reset(context.Background())
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New(("username expected for register command"))
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	s.config.SetUser(user.Name)
	fmt.Printf("User has been created: %v", user)

	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username expected for login command")
	}

	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	_, err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
