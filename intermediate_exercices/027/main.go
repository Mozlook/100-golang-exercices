// Exercise: SLICES

// Exercise: From the slice `myslice`, create a new slice containing only the middle 3 elements.
// Hint: use slice[start:end] syntax

package main

import "fmt"

func main() {
	slice := []int32{0, 1, 2, 3, 4}
	// Your code goes here
	newSlice := slice[1:4]
	fmt.Println(newSlice)
}

