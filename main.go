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

	// Start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1) 
        go worker(w, jobs, results) // worker runs in the background doing the scraping
    }
    // Send jobs (URLs) to workers
    go func() {
        for _, url := range urls {
            jobs <- url // push each url into channel for workers to grab
        }
        close(jobs) // no more urls after this
    }()

    // Collect results after all workers done
    go func() {
        wg.Wait()      // wait for ALL workers to finish
        close(results) // after that, we're done receiving results
    }()

    // Loop through results workers sent us
    for r := range results {
        if r.Error != nil {
            fmt.Printf("%s | ERROR: %v\n", r.URL, r.Error) // if any URL failed
        } else {
            fmt.Printf("%s | Status: %d | Size: %d bytes\n",
                r.URL, r.StatusCode, r.Size) // success: show url, status code, and size
        }
    }

	fmt.Println("\nScraping complete!")
}
