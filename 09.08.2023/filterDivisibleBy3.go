package main

import "fmt"

func filterNumbers(numbers []int, applyFilter func(int) bool) []int {
	var filterednumbers []int
	for _, num := range numbers {
		if applyFilter(num) {
			filterednumbers = append(filterednumbers, num)
		}
	}
	return filterednumbers
}

func main() {
	numbersSlice := []int{1, 2, 3, 12, 15, 14, 21, 18}

	divisibleBy3 := func(num int) bool {
		return num%3 == 0
	}

	filteredDivisibleBy3Slice := filterNumbers(numbersSlice, divisibleBy3)
	fmt.Println("Divisible by 3 elements are: ", filteredDivisibleBy3Slice)

}
