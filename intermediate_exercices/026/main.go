// Exercise: SLICES

// Create a slice of the first 5 numbers from a list of 10 numbers

package main

import "fmt"

func main() {
	myset := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Your code goes here
	newSlice := myset[0:5]
	fmt.Println(newSlice)
}

