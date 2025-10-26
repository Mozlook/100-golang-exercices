// Exercise: Channels - Range (unbuffered)

// In this exercise we will use the range keyword to iterate over a UNbuffered (sync) channel.
// Create a unbuffered channel (type int)
// Create a goroutine that:
// Has an infinite loop -> prints the current second and waits for a second
// Use the range iterator in the main function to see each second

package main

import (
	"fmt"
	"time"
)

func second(c chan int, multiplier int) {
	timer := 1
	for {
		c <- timer * multiplier
		time.Sleep(1 * time.Second)
		timer += multiplier
	}
}

func main() {
	var c chan int = make(chan int)
	go second(c, 2)
	go second(c, 3)
	for value := range c {
		fmt.Println(value)
	}
}
