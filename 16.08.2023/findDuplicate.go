/*remove the duplicate elements from array */

package main

import "fmt"

func removeDuplicate(arr []int) []int {
	slice := []int{}

	seen := make(map[int]bool)

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			slice = append(slice, num)
		}
	}
	return slice
}

func main() {
	arrayOfElements := []int{1, 2, 3, 4, 5, 2, 6, 7, 3, 8, 5}
	arrayWithNoDuplicates := removeDuplicate(arrayOfElements)

	fmt.Println("removed duplicate elements from array: ", arrayWithNoDuplicates)

}
