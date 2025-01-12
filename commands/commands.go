package commands

import (
	"errors"

	"github.com/emield/gator/internal/types"
)

type Commands struct {
	CommandMap map[string]func(*types.State, types.Command) error
}

func (c *Commands) Register(name string, f func(*types.State, types.Command) error) {
	c.CommandMap[name] = f
}

func (c *Commands) Run(s *types.State, cmd types.Command) error {
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
