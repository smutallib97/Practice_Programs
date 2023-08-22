package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)
	for i := 0; i <= 5; i++ {
		ch <- i

	}
	wg.Done()
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Println("Recieved:", num)

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
