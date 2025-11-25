package cmds

import (
	"context"
	"fmt"
	"log"
	"time"
)

func CmdFetch(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time ex. 1s 1m 1h>", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]
	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	log.Printf("Collecting feeds every %s...", duration)

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't fetch next feed", err)
		return
	}
	_, err = s.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed as fetched %w", err)
		return
	}
	feedData, err := s.ApiClient.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't get feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
