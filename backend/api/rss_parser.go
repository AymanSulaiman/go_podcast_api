package api

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Rss represents the root of an RSS feed
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
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

// Item represents an individual item in the podcast feed
type Item struct {
	Title          string         `xml:"title"`
	Link           string         `xml:"link"`
	Description    string         `xml:"description"`
	PubDate        string         `xml:"pubDate"`
	GUID           string         `xml:"guid"`
	Enclosure      Enclosure      `xml:"enclosure"`
	ContentEncoded string         `xml:"content:encoded"`
	Itunes         ItunesElements `xml:"itunes"` // Custom field for iTunes elements
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

// Enclosure represents the media enclosure in the feed
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

// FetchRssRequest represents the request structure for fetching an RSS feed
type FetchRssRequest struct {
	URL string `json:"url"`
}

// UnmarshalXML handles namespaced elements for an Item
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

// FetchRss fetches and parses an RSS feed from the given URL
func FetchRss(url string) (Rss, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Rss{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Rss{}, fmt.Errorf("non-200 status code received")
	}

	var rss Rss
	decoder := xml.NewDecoder(resp.Body)
	if err := decoder.Decode(&rss); err != nil {
		return Rss{}, err
	}

	return rss, nil
}
