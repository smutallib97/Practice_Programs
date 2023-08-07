/*
Write a Go function called FindMax that takes a slice of integers as input and returns
the maximum value from the slice.
*/
package main

import (
	"fmt"
)

func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		panic("Slice is empty")
	}

	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	numbers := []int{12, 56, 89, 7, 23, 45}
	max := FindMax(numbers)
	fmt.Printf("Maximum value: %d\n", max)
}
