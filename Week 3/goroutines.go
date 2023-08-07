/*
Write a Go program that creates two Go routines, each printing numbers from 1 to 5.
Ensure that the output from both Go routines is interleaved.
*/

package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup, ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Routine %d: %d\n", <-ch, i)
		ch <- 0
	}
	wg.Done()
}

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go printNumbers(wg, ch)
	go printNumbers(wg, ch)

	ch <- 0 // Start the first goroutine

	wg.Wait()
	close(ch)
}
