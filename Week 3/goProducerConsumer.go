/*
Design a Go program that simulates a simple producer-consumer scenario.
Create two Go routines: one to produce integers from 1 to 10 and send them to a channel,
and another to read these numbers from the channel and print them.
*/

package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
	wg.Done()
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go producer(ch)
	go consumer(&wg, ch)

	wg.Wait()
}
