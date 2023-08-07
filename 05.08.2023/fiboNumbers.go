//Implement a function to find the first N Fibonacci numbers.

package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	fibNumbers := make([]int, n)
	fibNumbers[0] = 0
	if n > 1 {
		fibNumbers[1] = 1
	}

	for i := 2; i < n; i++ {
		fibNumbers[i] = fibNumbers[i-1] + fibNumbers[i-2]
	}

	return fibNumbers
}

func main() {
	n := 10
	fibNumbers := fibonacci(n)
	fmt.Printf("First %d Fibonacci numbers: %v\n", n, fibNumbers)
}
