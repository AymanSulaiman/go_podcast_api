package main

import (
	"net/http" // From Go

	"github.com/Podcherry/podcherry_webapp/backend/api" // Local

	"github.com/gin-contrib/cors" // Import the CORS package
	"github.com/gin-gonic/gin"    // This makes life so much easier
)

//TODO
// Add in swagger for api documentation

func main() {
	r := gin.Default()

	// Use the CORS middleware
	r.Use(cors.Default()) // This will allow all origins by default

	r.GET("/shows", func(c *gin.Context) {
		term := c.Query("term") // example: /search?term=tech
		dataCh := make(chan []byte)
		errCh := make(chan error)

		go func() {
			data, err := api.SearchPodcastShow(term)
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
			episodes, err := api.SearchPodcastEpisodes(term)
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

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
