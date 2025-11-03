// Exercise: Set up a simple HTTP Server

// We will use the net/http library
// https://pkg.go.dev/net/http

// Like with ExpressJS, we will make a web server, and the web server will serve the "/bar" route
// and the response should be "Hello, /var". BUT don't hardcode the URI!
// Changing the http.HandleFunc 1st variable (the URI), the message served should also change.
// Example: If I set my webserver at /newpath, my response will be "Hello, /newpath"

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		response := fmt.Sprintf("Hello, %s", html.EscapeString(path))
		fmt.Fprintf(w, response)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

