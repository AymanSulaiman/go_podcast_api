package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SearchPodcastShow(term string) ([]byte, error) {
	encodedTerm := url.QueryEscape(term)
	url := "https://itunes.apple.com/search?term=" + encodedTerm + "&entity=podcast"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func SearchPodcastEpisodes(term string) ([]map[string]interface{}, error) {
	encodedTerm := url.QueryEscape(term)
	url := "https://itunes.apple.com/search?term=" + encodedTerm + "&entity=podcastEpisode"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Results []map[string]interface{} `json:"results"`
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
