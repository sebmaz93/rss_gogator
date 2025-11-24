package cmds

import (
	"context"
	"fmt"
)

func CmdFeeds(s *State, cmd Command) error {
	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	usersMap := make(map[string]string)
	for _, feed := range feeds {
		_, ok := usersMap[feed.UserID.String()]
		if !ok {
			user, err := s.DB.GetUserById(context.Background(), feed.UserID)
			if err != nil {
				return err
			}
			usersMap[feed.UserID.String()] = user.Name
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(usersMap[feed.UserID.String()])
	}

	return nil
}
