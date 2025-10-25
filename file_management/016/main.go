// Exercise: Check if a file exists
// Tip: use the "os" package

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Stat("./READM.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(file)
}

