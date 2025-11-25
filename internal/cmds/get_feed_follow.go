package cmds

import (
	"context"
	"fmt"
)

func GetFeedFollowsForUser(s *State, cmd Command) error {
	currUsername := s.Cfg.CurrentUserName
	user, err := s.DB.GetUser(context.Background(), currUsername)
	if err != nil {
		return err
	}
	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return nil
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
