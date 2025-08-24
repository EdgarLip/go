package examplesOnly

import (
	"log"
	"net/http"
	"time"
)

// this program is a nice example how can i send many cunccurent http requests and read data
// simultaniously.

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func RunGetUrls() {
	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}

	/*
	What happens here:
	1. go get(url, results) starts a new goroutine
	2. The goroutine runs in the background
	3. The main goroutine immediately continues to the next iteration
	4. No blocking occurs - each go get() call returns instantly

	Timeline Visualization:
	Main Goroutine:          Background Goroutines:
	  go get(amazon.com) ────► [Goroutine 1: fetching amazon.com...]
	  go get(google.com) ────► [Goroutine 2: fetching google.com...]
	  go get(nytimes.com) ───► [Goroutine 3: fetching nytimes.com...]
	  go get(wsj.com) ───────► [Goroutine 4: fetching wsj.com...]

	  First loop completes immediately!

	  Now waiting in second loop:
	  r := <-results ←──────── [Goroutine 2 completes first, sends result]
	  r := <-results ←──────── [Goroutine 1 completes, sends result]
	  r := <-results ←──────── [Goroutine 4 completes, sends result]
	  r := <-results ←──────── [Goroutine 3 completes, sends result]
	*/

	for _, url := range list {
		go get(url, results)
	}

	/*
	Why The Second Loop Works:
	  Each goroutine eventually calls: ch <- result{url, nil, t}
	  The channel receives 4 results (one from each goroutine)
	  The main goroutine reads 4 results (one per loop iteration)

	  You'll see that "All goroutines started..." prints immediately, then results
	  come back in random order as the HTTP requests complete.
	  The magic: The go keyword makes the function call non-blocking, allowing concurrent execution! 🚀
	*/

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}
