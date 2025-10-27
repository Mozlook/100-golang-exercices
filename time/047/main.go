// Exercise: Time

// Add 20 minutes to the current time (shown in UTC)

package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now().UTC()
	inTenMinutes := current.Add(10 * time.Minute)

	fmt.Println(inTenMinutes)
}
