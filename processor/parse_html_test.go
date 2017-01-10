package processor

import (
	"testing"
)

func TestCollectLinks(t *testing.T) {

	CollectLinks("https://news.ycombinator.com")
}
