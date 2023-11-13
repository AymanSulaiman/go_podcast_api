package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchPodcastShow_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonData := `{"results": [{"title": "Test Show"}]}`
		w.Write([]byte(jsonData))
	}))
	defer ts.Close()

	data, err := SearchPodcastShow("test")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var result struct {
		Results []map[string]interface{} `json:"results"`
	}
	json.Unmarshal(data, &result)

	if len(result.Results) != 1 || result.Results[0]["title"] != "Test Show" {
		t.Errorf("Expected 'Test Show', got '%v'", result.Results)
	}
}

func TestSearchPodcastShow_FailOnHTTPError(t *testing.T) {
	ts := httptest.NewServer(http.NotFoundHandler())
	defer ts.Close()

	_, err := SearchPodcastShow("test")
	if err == nil {
		t.Error("Expected an error, got none")
	}
}

// Similar test cases can be written for SearchPodcastEpisodes
// Consider testing for empty results, malformed JSON, etc.
