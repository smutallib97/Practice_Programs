//Write a function to calculate the factorial of a given number.

package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {

	var num1 int

	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num1)

	fmt.Println("Factorial of ", num1, "is :", factorial(num1))

}
