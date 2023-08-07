//Given a list of integers, find the maximum and minimum numbers.

package main

import "fmt"

func findMaxMin(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return max, min
}

func main() {
	numbers := []int{10, 5, 8, 2, 20, 15}
	max, min := findMaxMin(numbers)
	fmt.Printf("Maximum number: %d\n", max)
	fmt.Printf("Minimum number: %d\n", min)
}
