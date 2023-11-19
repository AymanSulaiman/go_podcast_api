# Podcast Search API
This API provides endpoints to search for podcasts and podcast episodes using the iTunes Search API.

## Setup
To run the server, make sure you have both Go and the Gin package installed.

```bash
go get -u github.com/gin-gonic/gin
```
Then, run the main file:

```bash
go run <filename>.go
```

The server will start on port 8080.

## Endpoints
### 1. Home
* URL: `/`
* Method: GET
* Response: A simple "Hello, World!" string.

### 2. Search Podcast Shows
* URL: `/shows?term=<SEARCH_TERM>`
* Method: GET
* Query Parameters:
    * term: The search term for the podcast. (e.g., tech)
* Response: A JSON response with the search results from the iTunes API.

### 3. Search Podcast Episodes
URL: /episodes?term=<SEARCH_TERM>
Method: GET
Query Parameters:
term: The search term for the podcast episodes. (e.g., tech news)
Response: A JSON response with the search results from the iTunes API.


## Implementation Details
### Functions
* `searchPodcastShow(term string) ([]byte, error)`:
    * Accepts a search term as its parameter.
    * Uses the iTunes Search API to find podcasts matching the search term.
    * Returns the raw JSON response as bytes.

* `searchPodcastEpisodes(term string) ([]map[string]interface{}, error)`:
    * Accepts a search term as its parameter.
    * Uses the iTunes Search API to find podcast episodes matching the search term.
    * Parses the JSON response to extract the episode details.
    * Returns a slice of map containing episode details.

### Concurrency
For improved performance, the API utilizes goroutines to concurrently search for podcast shows and episodes. Results or errors from these goroutines are communicated back to the main routine using channels.

## Building a Docker Container to test changes
````bash
docker build -t backend .
docker run -d -p 8000:8000 backend
````