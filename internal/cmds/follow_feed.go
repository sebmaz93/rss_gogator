package cmds

import (
	"context"
	"fmt"

	"github.com/sebmaz93/rss_gogator/internal/database"
)

func CmdFollowFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage %s <url>", cmd.Name)
	}

	currUsername := s.Cfg.CurrentUserName
	user, err := s.DB.GetUser(context.Background(), currUsername)
	if err != nil {
		return err
	}
	url := cmd.Args[0]
	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	feed_follow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	fmt.Println(feed_follow.UserName)
	fmt.Println(feed_follow.FeedName)

	return nil
}
