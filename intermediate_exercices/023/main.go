// Exercise: Random
// Generate a random number from te range [-50, +50]
package main

import (
	"fmt"
	"math/rand"
)

func random(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	randomNum := random(-50, +50)
	fmt.Printf("Random number: %d\n", randomNum)
}

