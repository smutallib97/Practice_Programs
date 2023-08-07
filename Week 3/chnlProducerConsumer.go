/*
How does Go's channel-based communication facilitate safe and synchronized data exchange between goroutines?
Provide an example of using channels to implement a producer-consumer pattern.
*/

package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	wg.Done()
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
	wg.Done()
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
