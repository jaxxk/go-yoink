package worker

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jaxxk/go-yoink/internal/database"
)

// RSSFEED represents the top-level RSS structure
type RSSFEED struct {
	XMLName xml.Name   `xml:"rss"`
	Channel rssChannel `xml:"channel"`
}

// rssChannel represents the <channel> element
type rssChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	Items       []rssItem `xml:"item"`
}

// CustomTime is a custom time type to handle the PubDate parsing
type CustomTime struct {
	time.Time
}

const customTimeFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

// UnmarshalXML is a custom unmarshaler for the CustomTime type
func (ct *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	// Parse the time according to the custom format
	t, err := time.Parse(customTimeFormat, v)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// rssItem represents the <item> element
type rssItem struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	PubDate     CustomTime `xml:"pubDate"`
	GUID        string     `xml:"guid"`
	Description string     `xml:"description"`
}

func FetchRSSFeed(url string) (*RSSFEED, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSSFEED
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return nil, err
	}

	return &rss, nil
}

func StartWorking(db *database.Queries, fetchFeed func(url string) (*RSSFEED, error)) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	ctx := context.Background()
	for {
		<-ticker.C

		// Step 1: Get the next `n` feeds to fetch
		feeds, err := db.GetNextFeedsToFetch(ctx, 10)
		if err != nil {
			log.Println("Error fetching feeds:", err)
			continue
		}

		var wg sync.WaitGroup
		for _, feed := range feeds {
			wg.Add(1)
			go func(feed database.GetNextFeedsToFetchRow) { // Use a goroutine to fetch the feed
				defer wg.Done()
				rssFeed, err := fetchFeed(feed.Url)
				if err != nil {
					log.Println("Error fetching feed:", feed.Url, err)
					return
				}
				// Process the feed (e.g., print titles)
				for _, item := range rssFeed.Channel.Items {
					err = processFeed(db, item, ctx, feed.ID)
					if err != nil {
						log.Printf("Error processing feeds item %s:", item.Title)
						continue
					}
				}
				markFeed := database.MarkFeedFetchedParams{
					UpdatedAt: time.Now(),
					ID:        feed.ID,
				}
				// Step 3: Mark feed as fetched
				err = db.MarkFeedFetched(ctx, markFeed)
				if err != nil {
					log.Println("Error marking feed as fetched:", markFeed, err)
				}
			}(feed)
		}
		wg.Wait() // Wait for all
	}
}

func processFeed(db *database.Queries, item rssItem, ctx context.Context, feedID string) error {
	post := database.CreatePostParams{
		ID:          uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       truncateString(item.Title, 128),
		Url:         truncateString(item.Link, 128),
		Description: truncateString(item.Description, 800),
		PublishedAt: item.PubDate.Time,
		FeedID:      feedID,
	}
	_, err := db.CreatePost(ctx, post)
	if err != nil {
		log.Println("Error creating post: ", err)
	}
	return nil
}
func truncateString(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[:maxLen]
	}
	return str
}
