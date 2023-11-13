package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func readJSONFromFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// func TestSearchPodcastShow_Success(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		jsonData := `{"results": [{"title": "Test Show"}]}`
// 		w.Write([]byte(jsonData))
// 	}))
// 	defer ts.Close()

// 	data, err := SearchPodcastShow("test")
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}

// 	var result struct {
// 		Results []map[string]interface{} `json:"results"`
// 	}
// 	json.Unmarshal(data, &result)

// 	if len(result.Results) != 1 || result.Results[0]["title"] != "Test Show" {
// 		t.Errorf("Expected 'Test Show', got '%v'", result.Results)
// 	}
// }

// func TestSearchPodcastShow_FailOnHTTPError(t *testing.T) {
// 	ts := httptest.NewServer(http.NotFoundHandler())
// 	defer ts.Close()

// 	_, err := SearchPodcastShow("test")
// 	if err == nil {
// 		t.Error("Expected an error, got none")
// 	}
// }

// // Similar test cases can be written for SearchPodcastEpisodes
// // Consider testing for empty results, malformed JSON, etc.

func TestSearchPodcastEpisodes(t *testing.T) {
	mockResponse, err := readJSONFromFile("test_episodes_search.json")
	if err != nil {
		t.Fatalf("Failed to read mock response file: %v", err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockResponse)
	}))
	defer ts.Close()

	episodes, err := SearchPodcastEpisodes("huberman")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(episodes) != 50 {
		t.Errorf("Expected 1 episode, got %d", len(episodes))
	}

	// Further assertions can be added to check specific fields in the response
}

func TestSearchPodcastShow_Huberman(t *testing.T) {
	mockResponse, err := readJSONFromFile("test_shows_search.json")
	if err != nil {
		t.Fatalf("Failed to read mock response file: %v", err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockResponse)
	}))
	defer ts.Close()

	data, err := SearchPodcastShow("huberman")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var result struct {
		ResultCount int                      `json:"resultCount"`
		Results     []map[string]interface{} `json:"results"`
	}
	json.Unmarshal(data, &result)

	if result.ResultCount != 6 {
		t.Errorf("Expected result count 6, got %d", result.ResultCount)
	}

	// Additional assertions can be made based on the structure of the results
}
