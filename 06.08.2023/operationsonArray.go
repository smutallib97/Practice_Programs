//Write a Go program that takes an array of integers as input.
//The program should calculate and print the sum, average, minimum, and maximum value of the array.

package main

import "fmt"

func calculateSum(input []int) int {
	sum := 0
	for _, numbers := range input {
		sum += numbers
	}
	return sum
}

func calculateAverage(input []int) float64 {
	sum := calculateSum(input)
	return float64(sum) / float64(len(input))
}

func findMinimum(input []int) int {
	min := input[0]

	for _, num := range input {
		if num < min {
			min = num
		}
	}
	return min
}
func findMaximum(input []int) int {
	max := input[0]

	for _, num := range input {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {

	inputArray := []int{5, 8, 10, 46, 3, 14}

	fmt.Println("Array of Integer numbers:", inputArray)

	calculateSumOfArrayElements := calculateSum(inputArray)
	fmt.Println("Sum of Array Elements is: ", calculateSumOfArrayElements)

	calculateAverageOfArray := calculateAverage(inputArray)
	fmt.Println("Average of Array is: ", calculateAverageOfArray)

	minimumValueInArray := findMinimum(inputArray)
	fmt.Println("Minimum value from Array is: ", minimumValueInArray)

	maximumValueOfArray := findMaximum(inputArray)
	fmt.Println("Maximum value of Array is: ", maximumValueOfArray)
}
