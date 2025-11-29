package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func mains() {
	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}

	const batchSize = 10
	var wg sync.WaitGroup

	// Create a context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	results := make(chan Result, len(urls))
	jobs := make(chan string)
	wg.Add(batchSize)

	for i := 0; i < batchSize; i++ {
		go func() {
			defer wg.Done()
			for url := range jobs {
				healthCheck(ctx, url, results)
			}

		}()
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		for res := range results {
			printer(res)
		}
	}()

	wg.Wait()
	close(results)

	println("All goroutines completed!")

	if ctx.Err() != nil {
		fmt.Println("\nTimeout occurred!")
	} else {
		fmt.Println("\nAll health checks completed!")
	}

}
