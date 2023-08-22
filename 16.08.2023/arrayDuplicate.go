package main

import "fmt"

// func findDuplicate(numbers []int, duplicates func(int) bool) []int {
// 	foundDuplicates := []int{}

// 	for _, n := range numbers {
// 		if duplicates(n) {
// 			foundDuplicates = append(foundDuplicates, n)
// 		}
// 	}
// 	return foundDuplicates
// }

func main() {

	array := []int{1, 2, 2, 3, 4, 5, 5, 6}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if array[j] == array[i] && i != j {
				fmt.Println(array[i])
			}
		}
	}

}
