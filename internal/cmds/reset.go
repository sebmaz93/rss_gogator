package cmds

import (
	"context"
	"fmt"
)

func CmdReset(s *State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error reseting users %w", err)
	}

	return nil
}
