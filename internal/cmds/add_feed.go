package cmds

import (
	"context"
	"fmt"

	"github.com/sebmaz93/rss_gogator/internal/database"
)

func CmdAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usave: %s <name> <url>", cmd.Name)
	}

	currUserName := s.Cfg.CurrentUserName
	user, err := s.DB.GetUser(context.Background(), currUserName)
	if err != nil {
		return err
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.ApiClient.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	fmt.Println(feed)

	return nil
}
