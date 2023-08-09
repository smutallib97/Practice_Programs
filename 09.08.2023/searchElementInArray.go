/*
Write a program to search an element in an array.
*/

package main

import "fmt"

func linearSearch(arr []int, numElement int) int {
	for i, number := range arr {
		if number == numElement {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{18, 12, 7, 2, 4, 19, 36}
	fmt.Println("Array :", arr)

	var n int
	fmt.Print("Enter a Element from above array to search: ")
	fmt.Scan(&n)
	fmt.Printf("\nSearching for %d\n", n)

	index := linearSearch(arr, n)
	if index != -1 {
		fmt.Printf("\n%d found at index %d \n", n, index)
	} else {
		fmt.Printf("\n%d not found in the array", n)
	}
}
