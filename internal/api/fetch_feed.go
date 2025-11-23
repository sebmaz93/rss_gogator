package api

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("user-agent", "gator")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(b, &rssFeed)
	if err != nil {
		return nil, err
	}
	decodeEscaped(&rssFeed)
	return &rssFeed, nil
}

func decodeEscaped(f *RSSFeed) {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	for i, v := range f.Channel.Item {
		f.Channel.Item[i].Title = html.UnescapeString(v.Title)
		f.Channel.Item[i].Description = html.UnescapeString(v.Description)
	}
}
