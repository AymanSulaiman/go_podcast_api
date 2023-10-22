package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/shows", func(c *gin.Context) {
		term := c.Query("term") // example: /search?term=tech
		dataCh := make(chan []byte)
		errCh := make(chan error)

		go func() {
			data, err := searchPodcastShow(term)
			if err != nil {
				errCh <- err
				return
			}
			dataCh <- data
		}()

		select {
		case data := <-dataCh:
			c.Data(http.StatusOK, "application/json", data)
		case err := <-errCh:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.GET("/episodes", func(c *gin.Context) {
		term := c.Query("term") // example: /episodes?term=tech news
		episodesCh := make(chan []map[string]interface{})
		errCh := make(chan error)

		go func() {
			episodes, err := searchPodcastEpisodes(term)
			if err != nil {
				errCh <- err
				return
			}
			episodesCh <- episodes
		}()

		select {
		case episodes := <-episodesCh:
			c.JSON(http.StatusOK, episodes)
		case err := <-errCh:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func searchPodcastShow(term string) ([]byte, error) {
	// URL-encode the search term
	encodedTerm := url.QueryEscape(term)

	// Using the iTunes Search API for podcasts with the encoded term
	url := "https://itunes.apple.com/search?term=" + encodedTerm + "&entity=podcast"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func searchPodcastEpisodes(term string) ([]map[string]interface{}, error) {
	// URL-encode the search term
	encodedTerm := url.QueryEscape(term)

	// Using the iTunes Search API for podcast episodes with the encoded term
	url := "https://itunes.apple.com/search?term=" + encodedTerm + "&entity=podcastEpisode"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Results []map[string]interface{} `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
