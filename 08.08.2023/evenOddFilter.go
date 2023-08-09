package main

import "fmt"

func filterNumbers(numbers []int, filter func(num int) bool) []int {
	var filteredNumbers []int
	for _, num := range numbers {
		if filter(num) {
			filteredNumbers = append(filteredNumbers, num)
		}
	}
	return filteredNumbers
}
func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenFilteredNumbers := func(num int) bool {
		return num%2 == 0
	}
	evenNumbers := filterNumbers(slice, evenFilteredNumbers)
	fmt.Println("Even Numbers Slice :", evenNumbers)

	oddFilteredSlice := func(num int) bool {
		return num%3 == 0
	}
	oddNumbers := filterNumbers(slice, oddFilteredSlice)
	fmt.Println("Odd Numbers Slice: ", oddNumbers)
}
