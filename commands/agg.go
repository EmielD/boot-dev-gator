package commands

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/emield/gator/internal/database"
	"github.com/emield/gator/internal/types"
	"github.com/google/uuid"
)

func HandlerAgg(s *types.State, cmd types.Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("agg usage: agg <time between fetches> - example: agg 1m")
	}

	duration, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", duration.String())
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		fmt.Println("Timer ran out, beginning to scrape feeds..")
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *types.State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	s.Db.MarkFeedFetched(context.Background(), feed.ID)

	for _, item := range rssFeed.Channel.Item {

		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
	}

	return nil
}

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, err
	}

	channel := &rssFeed.Channel
	channel.Title = html.UnescapeString(channel.Title)
	channel.Description = html.UnescapeString(channel.Description)

	for i := range channel.Item {
		item := &channel.Item[i]

		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}
