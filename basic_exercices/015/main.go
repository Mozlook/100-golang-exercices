// Exercise: User input
// Get a number from the console and check if it's odd
// You can create another function or do it everything in the main func :)

package main

import "fmt"

func main() {
	// Here goes your code
	var number int
	fmt.Println("Choose a number")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("That's not a valid number!")
		return
	}
	if number%2 == 0 {
		fmt.Println("Your number is even")
	} else {
		fmt.Println("Your number is odd")
	}
}

