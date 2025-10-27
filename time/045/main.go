// Exercise: Tickers

// Create a goroutine with a infinite loop
// The condition for that infinite loop will be a range that goes over the "time.Tick(time.Second * 1)", in other words: a simple ticker
// And at every Tick, we will print "Tick"
// In the main function, call the go routine :)

package main

import (
	"fmt"
	"time"
)

func task() {
	for tick := range time.Tick(time.Second) {
		fmt.Println("Tick:", tick)
	}
}

func main() {
	go task()

	time.Sleep(time.Second * 5)
}
