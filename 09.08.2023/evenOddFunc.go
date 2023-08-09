package main

import "fmt"

func numbersFilter(slice []int, filter func(int) bool) []int {
	var filteredNumber []int

	for _, num := range slice {
		if filter(num) {
			filteredNumber = append(filteredNumber, num)
		}
	}
	return filteredNumber
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Slice is: ", numbers)
	evenNumbers := func(n int) bool {
		return n%2 == 0
	}
	filteredEvenNumbers := numbersFilter(numbers, evenNumbers)
	fmt.Println("Even Numbers Slice is:", filteredEvenNumbers)

}
