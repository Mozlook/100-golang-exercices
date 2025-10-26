// Exercise: Channels

// It's important to get this concept, document yourself first! :)
// There are two concurrent routines (f1 and f2)
// Send a message "Hello from f1" from f1 function
// Receive the message from f1 into the f2 function, and print "I am f2 and ..." + the message from f1
// This should be done with a channel

package main

import (
	"fmt"
	"time"
)

func f1(c chan string) {
	msg := "Hello from f1"
	c <- msg
}

func f2(c chan string) {
	msg := <-c
	fmt.Printf("I am f2 and %s", msg)
}

func main() {
	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	ch := make(chan string)
	go f1(ch)
	go f2(ch)
	time.Sleep(1 * time.Second)
}

