/*Write a Go function called calculateSum that takes an integer n as input and returns
the sum of all positive integers from 1 to n. For example, if n is 5,
the function should return 15 (1 + 2 + 3 + 4 + 5).*/

package main

import "fmt"

func calculateSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func main() {
	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil || num < 0 {
		fmt.Println("Invalid Input, enter non-negative integer")
		return
	}

	result := calculateSum(num)
	fmt.Printf("Sum of positive integers from 1 to %d is:  %d", num, result)
}
