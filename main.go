package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

// Structure to store results
type FetchResult struct {
	URL        string
	StatusCode int
	Size       int
	Error      error
}

// Worker function
func worker(id int, jobs <-chan string, results chan<- FetchResult) {
	defer wg.Done() // when everything is done

	for url := range jobs { // go through the urls
		resp, err := http.Get(url) // try to hit url with http get
		if err != nil {
			results <- FetchResult{URL: url, Error: err}
			continue
		}
		defer resp.Body.Close() // close after reading

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			results <- FetchResult{
				URL:        url,
				StatusCode: resp.StatusCode,
				Error:      err,
			}
			continue
		}

		results <- FetchResult{ // send back data (url, status, size)
			URL:        url,
			StatusCode: resp.StatusCode,
			Size:       len(body),
			Error:      nil,
		}
	}
}

func main() {

	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://uottawa.ca",
		"https://github.com",
		"https://httpbin.org/get",
	}

	numWorkers := 3

	jobs := make(chan string, len(urls))
	results := make(chan FetchResult, len(urls))

	// TODO: Start workers

	// TODO: Send jobs

	// TODO: Collect results

	fmt.Println("\nScraping complete!")
}
