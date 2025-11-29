package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
)

func main() {
	var g errgroup.Group
	g.SetLimit(3)

	for i := 1; i <= 10; i++ {
		i := i // Create a new variable for this iteration (Go < 1.22)
		g.Go(func() error {
			fmt.Printf("Goroutine %d is starting\n", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d is done\n", i)
			return nil
		})
	}

	fmt.Println("Goroutines before Wait():", runtime.NumGoroutine())

	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered an error: %v\n", err)
	}
	fmt.Println("All goroutines complete.")

	fmt.Println("Goroutines after Wait():", runtime.NumGoroutine())
}
