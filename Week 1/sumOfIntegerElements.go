/*Write a function in Go that takes a slice of integers and returns the sum of its elements.*/

package main

import (
	"fmt"
)

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("Sum of elements: %d\n", sum(numbers))
}
