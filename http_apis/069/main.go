// Exercise: JSON!! - Marshal (encode)

// We will create a json object.
// First, create a struct called 'user' which contains 2 attributes of type string called 'Username' and 'Password'
// Then, create an array of 3 users 'John', 'Anna' and 'Mike' with any random password
// Use the json.Marshal() function to encode the go array into a JSON object.

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	people := []user{
		{Username: "John", Password: "qwerty"},
		{Username: "Anna", Password: "qwerty"},
		{Username: "Mike", Password: "qwerty"},
	}

	out, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	// print encoded json data
	fmt.Println(string(out))
}

