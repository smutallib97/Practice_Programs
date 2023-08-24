/*Write a Go program that finds the missing number in an array of consecutive integers from 1 to n
where one number is missing.*/

package main

import "fmt"

func getMissingElement(arr []int, n int) int {

	total := (n + 1) * (n + 2) / 2
	for i := 0; i < n; i++ {
		total -= arr[i]
	}
	return total
}

func main() {
	arr := []int{1, 3, 4, 5, 6}
	N := len(arr)
	missingElement := getMissingElement(arr, N)

	fmt.Printf("Missing Element from the array is: %d", missingElement)

}
