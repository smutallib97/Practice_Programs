package main

import "fmt"

func main() {
	originalSlice := [5]int{1, 2, 3, 4, 5}

	duplicateSlice := [5]int{}

	for i := 0; i < len(originalSlice); i++ {
		duplicateSlice[i] = originalSlice[i]
	}
	fmt.Println("Original Slice: ", originalSlice)
	fmt.Println("Duplicate Slice: ", duplicateSlice)
}
