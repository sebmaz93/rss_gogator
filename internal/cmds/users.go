package cmds

import (
	"context"
	"fmt"
)

func CmdListUsers(s *State, cmd Command) error {
	users, err := s.DB.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching users %w", err)
	}
	for _, user := range users {
		if s.Cfg.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}

	}
	return nil
}
