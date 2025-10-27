// Exercise: Time

// We will have 2 time variables
// You will have to:
// 1- Get the difference with the Sub() function
// 2- See if they are equal with the Equal() funciton
// 3- See what comes after the other with the After() function

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2020, 2, 1, 3, 0, 0, 0, time.UTC)
	end := time.Date(2021, 2, 1, 12, 0, 0, 0, time.UTC)
	difference := end.Sub(start)
	fmt.Println("Difference:", difference)
	equal := start.Equal(end)
	fmt.Println("Is equal:", equal)
	if start.After(end) {
		fmt.Println(start, "is Later")
	} else {
		fmt.Println(end, "is Later")
	}
}
