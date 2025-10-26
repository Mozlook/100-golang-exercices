// Exercise: Random
// Create a program that mimics a dice roll (a 6-face dice)
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Your code goes here
	r := rand.Intn(6)
	fmt.Println(r + 1)
}

