/*
Write a Go program that uses goroutines to perform concurrent tasks.
Use mutexes and channels to coordinate access to a shared resource among multiple goroutines,
ensuring that only one goroutine can modify the resource at a time.
*/

package main

import (
	"fmt"
	"sync"
)

var (
	sharedResource int
	mutex          sync.Mutex
)

func modifySharedResource(ch chan<- bool) {
	mutex.Lock()
	defer mutex.Unlock()

	// Modify the shared resource
	sharedResource++
	fmt.Println("Shared resource modified by goroutine")

	ch <- true
}

func main() {
	numGoroutines := 5
	ch := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		go modifySharedResource(ch)
	}

	// Wait for all goroutines to finish
	for i := 0; i < numGoroutines; i++ {
		<-ch
	}

	fmt.Println("Final shared resource value:", sharedResource)
}
