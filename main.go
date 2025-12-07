package main

import (
	"fmt"
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
	// TODO: fetch the URL
	// TODO: send Result struct to results channel
	// hint: use resp, err := http.Get(url)
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
