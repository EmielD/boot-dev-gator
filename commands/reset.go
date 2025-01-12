package commands

import (
	"context"

	"github.com/emield/gator/internal/types"
)

func HandlerReset(s *types.State, cmd types.Command) error {
	s.Db.Reset(context.Background())
	return nil
}
