package main

import (
	"errors"
	"fmt"

	"github.com/emield/gator/internal/config"
)

type state struct {
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username expected for login command")
	}

	_, err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
