// Exercise: Read a file

// beware: You should run this code where the read file is, or reference it!
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	rawData, err := os.ReadFile("./read.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data string = string(rawData)
	fmt.Println(data)
}
