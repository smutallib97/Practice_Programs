/*Write a Go program that calculates the sum of the first n terms in the Fibonacci sequence.
The value of 'n' should be taken as input from the user.*/

package main

import "fmt"

func fibonacciSum(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		sum += a
		a, b = b, a+b
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Enter a positive integer number: ")
	fmt.Scan(&n)

	res := fibonacciSum(n)
	fmt.Printf("Sum of first %d terms in the fibonacci sequence is: %d", n, res)
}
