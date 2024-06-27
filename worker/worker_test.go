package worker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchRSSFeed(t *testing.T) {
	rssFeed, err := FetchRSSFeed("https://blog.boot.dev/index.xml")

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, rssFeed.Channel.Title, "Boot.dev Blog")
}
