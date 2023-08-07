/*
Create a simple concurrent web server in Go that serves multiple HTTP requests concurrently using Go routines.
The server should respond to each request with a simple "Hello, World!" message.
*/

package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
