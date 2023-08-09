/*
Write a go program to double the values of slice using Map and Higher Order function.
*/

package main

import "fmt"

func mapFunc(numbers []int, doubleFunc func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = doubleFunc(num)
	}
	return result
}

func doubleValue(x int) int {
	return x * 2
}

func main() {
	numbersSlice := []int{1, 2, 3, 4, 5}

	doubleNumbers := mapFunc(numbersSlice, doubleValue)

	fmt.Println(doubleNumbers)

}
