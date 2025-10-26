// Exercise: Channels directions (only read/rx)

// Make a goroutine with a channel for only receive data.
// The function should be called "receive" and the receive-only channel should be it's 1st and only argument
// Sending data from that channel is prohibited / will cause compiler errors
// Feed some string into that channel.

package main

import (
	"fmt"
	"time"
)

func receive(in <-chan string) {
	msg := <-in
	fmt.Print(msg)
}

func main() {
	c := make(chan string)
	go receive(c)
	c <- "test"
	time.Sleep(5 * time.Second)
}

