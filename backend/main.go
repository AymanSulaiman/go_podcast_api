package main

import (
	"net/http"

	"github.com/Podcherry/podcherry_webapp/backend/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO
// Add in swagger for api documentation
// Load Testing and high concurrancy, hard to do concurrancy from Isolation for root cause analysis

func main() {
	r := gin.Default()

	// Use the CORS middleware
	r.Use(cors.Default()) // This will allow all origins by default
	// config := cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080", "http://localhost:5000"}, // Or your Svelte app's production URL
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }
	// r.Use(cors.New(config))

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

	// Define the /fetch-rss endpoint
	r.POST("/fetch-rss", func(c *gin.Context) {
		var req api.FetchRssRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		rssCh := make(chan api.Rss) // Channel to receive the parsed RSS
		errCh := make(chan error)   // Channel to receive any errors

		// Use a goroutine to fetch and parse the RSS asynchronously
		go func() {
			rss, err := api.FetchRss(req.URL)
			if err != nil {
				errCh <- err
				return
			}
			rssCh <- rss
		}()

		// Use select to wait on multiple channel operations
		select {
		case rss := <-rssCh:
			c.JSON(http.StatusOK, rss)
		case err := <-errCh:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8000") // listen and serve on localhost:8000
}
