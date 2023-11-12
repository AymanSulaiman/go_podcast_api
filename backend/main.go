package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/Podcherry/podcherry_webapp/backend/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Rss represents the root of an RSS feed
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type FetchRssRequest struct {
	URL string `json:"url"`
}

// Channel contains details about the podcast channel
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Language    string `xml:"language"`
	Copyright   string `xml:"copyright"`
	Description string `xml:"description"`
	Image       Image  `xml:"image"`
	Items       []Item `xml:"item"`
}

// Image represents the podcast image details
type Image struct {
	URL   string `xml:"url"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Item struct {
	Title          string         `xml:"title"`
	Link           string         `xml:"link"`
	Description    string         `xml:"description"`
	PubDate        string         `xml:"pubDate"`
	GUID           string         `xml:"guid"`
	Enclosure      Enclosure      `xml:"enclosure"`
	ContentEncoded string         `xml:"content:encoded"`
	Itunes         ItunesElements // Custom field for iTunes elements
}

// ItunesElements holds namespaced iTunes elements
type ItunesElements struct {
	Title       string `xml:"title"`
	EpisodeType string `xml:"episodeType"`
	Episode     string `xml:"episode"`
	Author      string `xml:"author"`
	Subtitle    string `xml:"subtitle"`
	Summary     string `xml:"summary"`
	Duration    string `xml:"duration"`
	Explicit    string `xml:"explicit"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

// Custom UnmarshalXML to handle namespaced elements
func (i *Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Item // Prevents infinite loop

	aux := &struct {
		ItunesTitle       string `xml:"itunes:title"`
		ItunesEpisodeType string `xml:"itunes:episodeType"`
		ItunesEpisode     string `xml:"itunes:episode"`
		ItunesAuthor      string `xml:"itunes:author"`
		ItunesSubtitle    string `xml:"itunes:subtitle"`
		ItunesSummary     string `xml:"itunes:summary"`
		ItunesDuration    string `xml:"itunes:duration"`
		ItunesExplicit    string `xml:"itunes:explicit"`
		*Alias
	}{
		Alias: (*Alias)(i), // Points to the original struct, avoiding recursion
	}

	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}

	// Map the decoded elements to the Itunes struct
	i.Itunes = ItunesElements{
		Title:       aux.ItunesTitle,
		EpisodeType: aux.ItunesEpisodeType,
		Episode:     aux.ItunesEpisode,
		Author:      aux.ItunesAuthor,
		Subtitle:    aux.ItunesSubtitle,
		Summary:     aux.ItunesSummary,
		Duration:    aux.ItunesDuration,
		Explicit:    aux.ItunesExplicit,
	}

	return nil
}

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

	r.POST("/fetch-rss", func(c *gin.Context) {
		var req FetchRssRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		rssCh := make(chan Rss)
		errCh := make(chan error)

		go func() {
			resp, err := http.Get(req.URL)
			if err != nil {
				errCh <- err
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				errCh <- fmt.Errorf("non-200 status code received")
				return
			}

			var rss Rss
			decoder := xml.NewDecoder(resp.Body)
			if err := decoder.Decode(&rss); err != nil {
				errCh <- err
				return
			}

			rssCh <- rss
		}()

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
