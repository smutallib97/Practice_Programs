/*
Write a program to find the prime numbers between two given numbers.
*/

package main

import (
	"fmt"
	"math"
)

func isPrimeFunc(number int) bool {
	if number == 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	i := 5

	for i*i <= number {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
		i = +6
	}
	return true
}
func findPrimesBetweenRange(start, end int) []int {
	primeNumbers := []int{}
	for num := int(math.Max(2, float64(start))); num <= end; num++ {
		if isPrimeFunc(num) {
			primeNumbers = append(primeNumbers, num)
		}
	}
	return primeNumbers
}
func main() {

	var startRange, endRange int
	fmt.Print("Enter the starting number: ")
	fmt.Scan(&startRange)
	fmt.Print("Enter the ending number: ")
	fmt.Scan(&endRange)

	primeNumbers := findPrimesBetweenRange(startRange, endRange)

	fmt.Printf("Prime numbers between %d to %d are : %d", startRange, endRange, primeNumbers)

}
