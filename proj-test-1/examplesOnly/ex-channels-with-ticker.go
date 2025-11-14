package examplesOnly

import (
	"log"
	"time"
)

func TickerWithChannelsAndSelect() {
	log.Println("start")

	const tickRate = 2 * time.Second

	stopper := time.After(5 * tickRate)
	ticker := time.NewTicker(tickRate).C

// "loop" is a label used to provide a named target for the "break" statement inside the select block below.
// In Go, you can label loops and then use "break <label>" to exit out of that specific loop from within nested blocks (like select or switch).
// This is especially useful when you want to break out of the outer loop from inside a select or switch statement.
// In this example, when the stopper channel receives a value, "break loop" will immediately exit the for loop below.
loop:
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}
	log.Println("finish")
}
