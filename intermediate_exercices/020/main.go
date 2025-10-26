// Exercise: STRUCT
// Create a Hotel structure with:
// numRooms int32
// streetName string
// hasPool bool

// Then assign a value to each of those attributes

package main

import "fmt"

type Hotel struct {
	// Your code goes here
	numRooms   int32
	streetName string
	hasPool    bool
}

func main() {
	var myHotel Hotel
	// Your code goes here
	myHotel.numRooms = 10
	myHotel.streetName = "Grove street"
	myHotel.hasPool = false

	fmt.Println(myHotel)
}

