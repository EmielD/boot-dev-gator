package types

import (
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
