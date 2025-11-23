package cmds

import (
	"context"
	"fmt"

	"github.com/sebmaz93/rss_gogator/internal/api"
)

func CmdFetch(s *State, cmd Command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("usage: %s <name>", cmd.Name)
	// }

	url := "https://www.wagslane.dev/index.xml"
	feed, err := api.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
