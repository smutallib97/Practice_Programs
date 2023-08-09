/*
Write a program to sort an array in ascending order.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbersArray := []int{8, 5, 6, 4, 7, 3, 9, 2, 10, 1}
	fmt.Println("Original Array : ", numbersArray)
	sort.Ints(numbersArray)
	fmt.Println("After sorting given array in acending order: ", numbersArray)
}
