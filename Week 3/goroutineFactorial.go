/*
Create a Go program that calculates the factorial of a given number concurrently using Go routines.
Prompt the user to input the number, and then display the result.
*/

package main

import (
	"fmt"
	"sync"
)

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go factorial(n, ch)

	go func() {
		result := <-ch
		fmt.Printf("Factorial of %d is %d\n", n, result)
		wg.Done()
	}()

	wg.Wait()
	close(ch)
}
