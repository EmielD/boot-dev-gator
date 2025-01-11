package commands

import (
	"errors"

	"github.com/emield/gator/internal/config"
	"github.com/emield/gator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	CommandMap map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CommandMap[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	if s == nil {
		return errors.New("could not run command, state is nil")
	}

	var command = c.CommandMap[cmd.Name]
	if command == nil {
		return errors.New("command does not exist")
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}
