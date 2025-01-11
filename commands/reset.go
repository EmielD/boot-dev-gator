package commands

import "context"

func HandlerReset(s *State, cmd Command) error {
	s.Db.Reset(context.Background())
	return nil
}
