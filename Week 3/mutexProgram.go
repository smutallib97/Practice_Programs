/*
You are designing a multi-threaded application that involves three shared resources: A, B, and C.
Implement a deadlock avoidance strategy using mutexes and synchronization mechanisms to ensure that
the application can never enter a deadlock state.
*/

package main

import (
	"fmt"
	"sync"
)

type Resource struct {
	Name     string
	mutex    sync.Mutex
	assigned bool
}

var (
	resourceA Resource
	resourceB Resource
	resourceC Resource
)

func process1() {
	resourceA.mutex.Lock()
	defer resourceA.mutex.Unlock()

	fmt.Println("Process 1 acquired resource A")
	resourceA.assigned = true

	resourceB.mutex.Lock()
	defer resourceB.mutex.Unlock()

	fmt.Println("Process 1 acquired resource B")
	resourceB.assigned = true

	resourceC.mutex.Lock()
	defer resourceC.mutex.Unlock()

	fmt.Println("Process 1 acquired resource C")
	resourceC.assigned = true

	fmt.Println("Process 1 completed successfully!")
}

func process2() {
	resourceB.mutex.Lock()
	defer resourceB.mutex.Unlock()

	fmt.Println("Process 2 acquired resource B")
	resourceB.assigned = true

	resourceC.mutex.Lock()
	defer resourceC.mutex.Unlock()

	fmt.Println("Process 2 acquired resource C")
	resourceC.assigned = true

	resourceA.mutex.Lock()
	defer resourceA.mutex.Unlock()

	fmt.Println("Process 2 acquired resource A")
	resourceA.assigned = true

	fmt.Println("Process 2 completed successfully!")
}

func main() {
	resourceA.Name = "Resource A"
	resourceB.Name = "Resource B"
	resourceC.Name = "Resource C"

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		process1()
	}()

	go func() {
		defer wg.Done()
		process2()
	}()

	wg.Wait()
}
