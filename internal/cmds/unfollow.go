package cmds

import (
	"context"
	"fmt"

	"github.com/sebmaz93/rss_gogator/internal/database"
)

func CmdUnfollowFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <URL>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
