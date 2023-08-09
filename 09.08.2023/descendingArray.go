/*
Write a program to sort an array in descending order.
*/

package main

import "fmt"

func bubbleDescendingSort(arr []int) {
	n := len(arr)
	for i := 0; i <= n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	numbers := []int{11, 5, 8, 12, 15, 18, 4, 20}
	fmt.Println("Before sorting :", numbers)

	bubbleDescendingSort(numbers)
	fmt.Println("After descending sorting :", numbers)

}
