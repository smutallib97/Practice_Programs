//Implement a function to find the LCM (Least Common Multiple) of two numbers.

package main

import "fmt"

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonMultiple(a, b int) int {
	return (a * b) / greatestCommonDivisor(a, b)
}

func main() {

	num1, num2 := 12, 10

	res1 := leastCommonMultiple(num1, num2)

	fmt.Printf("Least Common Multiple of %d and %d is : %d", num1, num2, res1)

}
