package cmds

import (
	"context"
	"fmt"

	"github.com/sebmaz93/rss_gogator/internal/database"
)

func CmdRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		Name: username,
	})
	fmt.Printf("DB user created %v", user.Name)
	if err != nil {
		return fmt.Errorf("error registering user %w\n", err)
	}
	if err = s.Cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("User %v registered", user.Name)

	return nil
}
