// Exercise: Write a list of 5 countries to a file
// Tip: use the "os" package

package main

import "os"

func main() {
	// Here goes your code
	file, _ := os.Create("./something.md")

	file.WriteString("Poland\nFrance\nSpain\nUSA\nNigeria")
}

