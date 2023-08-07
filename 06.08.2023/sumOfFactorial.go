/*Write a Go program that calculates the sum of the factorials of all positive integers from 1 to n.
The value of 'n' should be taken as input from the user.*/

package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 0
	}
	res := 1
	for i := 1; i <= n; i++ {
		res += i
	}
	return res

}

func sumOfFactorial(num int) int {
	sum := 0

	for i := 1; i <= num; i++ {
		sum += factorial(i)
	}
	return sum
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil || number < 0 {
		return
	}

	result := sumOfFactorial(number)
	fmt.Printf("Sum of the factorials of all positive integers from 1 to %d is: %d", number, result)

}
