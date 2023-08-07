//Write a function to perform a binary search on a sorted list.

package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	numbers := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 9
	index := binarySearch(numbers, target)
	if index != -1 {
		fmt.Printf("%d is found at index %d\n", target, index)
	} else {
		fmt.Printf("%d is not found in the list\n", target)
	}
}
