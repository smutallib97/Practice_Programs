//Write a function to find the sum of all the even numbers in a given list.

package main

import "fmt"

func sumOfEvenNumbers(numbers []int) int {
	var sum int = 0

	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	numberlist := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := sumOfEvenNumbers(numberlist)

	fmt.Println("Sum of even numbers : ", result)
}
