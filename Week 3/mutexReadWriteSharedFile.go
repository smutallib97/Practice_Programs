/*
Design a concurrent program that simulates multiple processes reading and writing to a shared file.
Use mutexes and synchronization to ensure that processes do not interfere with each other,
reventing potential data corruption or deadlocks.
*/

package main

import (
	"fmt"
	"sync"
)

var (
	fileContent string
	mutex       sync.Mutex
)

func readFromSharedFile() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("Reading from shared file:", fileContent)
}

func writeToSharedFile(data string) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("Writing to shared file:", data)
	fileContent = data
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		readFromSharedFile()
	}()

	go func() {
		defer wg.Done()
		writeToSharedFile("Hello, World!")
	}()

	wg.Wait()
}
