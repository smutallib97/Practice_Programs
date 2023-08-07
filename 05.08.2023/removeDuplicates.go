//Write a function to remove duplicates from a list of integers.

package main

import "fmt"

func removeDuplicates(nums []int) []int {
	slice := []int{}
	seen := make(map[int]bool)

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			slice = append(slice, num)
		}
	}

	return slice
}

func main() {
	numbersSlice := []int{1, 2, 3, 2, 4, 5, 5, 6, 7, 6}
	uniqueNumbers := removeDuplicates(numbersSlice)
	fmt.Println("List with duplicates removed:", uniqueNumbers)
}
