package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestup23892183912idname.com/",
		"http://kasia.in",
	}

	wg.Add(len(urls))
	for _, url := range urls {
		// Increment the WaitGroup counter.
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			response, err := http.Get(url)
			if err != nil {
				fmt.Println("Error for url:", url, ":", err.Error())
			} else {
				fmt.Println("Result for", url, "is:", response.StatusCode)
			}
		}(url)
	}

	// Wait for all HTTP fetches to complete.
	wg.Wait()

	fmt.Printf("%+v\n", "DONE!")
}
