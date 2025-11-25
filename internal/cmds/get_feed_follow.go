package cmds

import (
	"context"
	"fmt"

	"github.com/sebmaz93/rss_gogator/internal/database"
)

func GetFeedFollowsForUser(s *State, cmd Command, user database.User) error {

	feeds, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return nil
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
