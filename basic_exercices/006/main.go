// Exercise: Arrays
// Create an array of 10 "int8" values, in it's initialization, fill those values from 0 to 9

package main

import "fmt"

func main() {
	// Here goes your code
	arrayInt8 := new([10]int)

	for i := 0; i < len(arrayInt8); i++ {
		arrayInt8[i] = i
	}
	fmt.Printf("%v", arrayInt8)
}

