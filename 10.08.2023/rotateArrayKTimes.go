package main

import (
	"fmt"
)

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
func rotateArray(arr []int, k int) {
	n := len(arr)
	k = k % n
	reverse(arr, 0, n-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	var k int
	fmt.Print("Enter How many times you want to rotate array: ")
	fmt.Scan(&k)
	fmt.Println("Original Array:", array)
	rotateArray(array, k)
	fmt.Println("Rotated Array:", array)
}
