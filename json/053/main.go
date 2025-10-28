// Exercise: JSON!! - Unknown beforehand data

// Now the JSON will be unstructured (we won't know the values beforehand)
// When data is unstructured, we will create a map of strings to empty interfaces.
// Create a map of string to an empty interface called "result"
// Unmarshall (decode) the json into the result
// Get the "birds" object into a variable called "birds" and print it's values

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	birdJson := `{
                "birds": {
                  "pigeon":"likes to perch on rocks",
                  "eagle":"bird of prey"
                },
                "animals": "none"
              }`
	var result map[string]interface{}

	json.Unmarshal([]byte(birdJson), &result)
	birds := result["birds"]
	fmt.Print(birds)
}

