//Given a list of integers, write a function to find the average of the numbers.

package main

import "fmt"

func averageOfIntegerNumbers(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func main() {
	listOfIntergerNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res1 := averageOfIntegerNumbers(listOfIntergerNumbers)

	fmt.Println("Average of interger numbers : ", res1)

}
