package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRss(t *testing.T) {
	// Example of a simple test
	// In real-world scenarios, consider using a mocking library to mock HTTP responses

	// Setting up a test server to mimic an RSS feed
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<rss><channel><title>Test Channel</title></channel></rss>`))
	}))
	defer ts.Close()

	rss, err := FetchRss(ts.URL)
	if err != nil {
		t.Errorf("FetchRss returned an unexpected error: %v", err)
	}

	if rss.Channel.Title != "Test Channel" {
		t.Errorf("Expected channel title 'Test Channel', got '%s'", rss.Channel.Title)
	}

	// Add more test cases to cover different scenarios, including error cases
}
