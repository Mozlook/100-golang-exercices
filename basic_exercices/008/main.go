// Create an array of 5 strings, and initialize it's 2 first values with some random names

package main

import "fmt"

func main() {
	arr := [5]string{"Golang", "Python"}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element index %d %v \n", i, arr[i])
	}
}

