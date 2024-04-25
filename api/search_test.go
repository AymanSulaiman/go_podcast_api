package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func readJSONFromFile(filePath string) ([]byte, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a buffer to hold the file contents
	var buffer bytes.Buffer

	// Use io.Copy to read the file content into the buffer
	if _, err := io.Copy(&buffer, file); err != nil {
		return nil, err
	}

	// Return the bytes from the buffer
	return buffer.Bytes(), nil
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

// 	_, err := SearchPodcastShow("testtesttestest")
// 	if err == nil {
// 		t.Error("Expected an error, got none")
// 	}
// }

func TestSearchPodcastShow_EmptySearchTerm(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("term") == "" {
			http.Error(w, "Bad Request: Empty Search Term", http.StatusBadRequest)
		}
	}))
	defer ts.Close()

	_, err := SearchPodcastShow(ts.URL + "?term=") // Simulate an empty search term
	if err == nil {
		t.Error("Expected an error for empty search term, got none")
	}
}

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

	// Construct the full URL to pass to the function
	fullURL := ts.URL + "?term=" + url.QueryEscape("huberman")

	data, err := SearchPodcastShow(fullURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var result struct {
		ResultCount int                      `json:"resultCount"`
		Results     []map[string]interface{} `json:"results"`
	}
	json.Unmarshal(data, &result)

	if result.ResultCount != 7 {
		t.Errorf("Expected result count 7, got %d", result.ResultCount)
	}

	// Additional assertions can be made based on the structure of the results
}
