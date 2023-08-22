/*Write a Go program to generate the Fibonacci sequence up to a given number n.
The Fibonacci sequence starts with 0 and 1, and each subsequent number is the sum of the two preceding ones.*/

package main

import "fmt"

func generateFibonacci(num int) []int {
	fibonacciSeq := []int{0, 1}

	for i := 2; ; i++ {
		nextFibonacci := fibonacciSeq[i-1] + fibonacciSeq[i-2]
		if nextFibonacci > num {
			break
		}
		fibonacciSeq = append(fibonacciSeq, nextFibonacci)
	}
	return fibonacciSeq
}

func main() {
	var num int

	fmt.Print("Enter a number:")
	fmt.Scan(&num)

	fibonacciSequence := generateFibonacci(num)
	fmt.Printf("Fibonacci numbers sequence upto %d : %v", num, fibonacciSequence)
}
