package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, there\n")
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
