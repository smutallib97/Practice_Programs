/*Write a Go program that prints all the prime numbers between 1 and 100 using a loop
and the appropriate control structures.*/

package main

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Prime numbers between 1 and 100:")

	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
