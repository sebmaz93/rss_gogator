package cmds

import (
	"context"
	"fmt"
)

func CmdLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]
	user, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error fetching user %w", err)
	}
	if err := s.Cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s has logged in", s.Cfg.CurrentUserName)
	return nil
}
