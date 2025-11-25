package middleware

import (
	"context"

	c "github.com/sebmaz93/rss_gogator/internal/cmds"
	db "github.com/sebmaz93/rss_gogator/internal/database"
)

func LoggedIn(handler func(s *c.State, cmd c.Command, user db.User) error) func(*c.State, c.Command) error {
	return func(s *c.State, cmd c.Command) error {
		user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
