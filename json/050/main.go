// Exercise: JSON!!

// Create a struct "Human", with 2 string attributes: Name and Description

// Keep the Human struct, now in the json object there are going to be 2 humans
// create a "humans" array of type Human and print it's values.

package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name        string
	Description string
}

func main() {
	humansJson := `[{"name": "Rick",
                 "description": "has a grandson called Morty"},
                {"name": "Cactus",
                 "description": "is one of the powerpuff girls' name"}]`

	var humans []Human
	err := json.Unmarshal([]byte(humansJson), &humans)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, person := range humans {
		fmt.Printf("Name %s, Description: %s\n", person.Name, person.Description)
	}
}
