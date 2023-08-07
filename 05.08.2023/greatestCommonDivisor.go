//    Implement a function to find the greatest common divisor (GCD) of two numbers.

package main

import "fmt"

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	num1, num2 := 12, 10

	res := greatestCommonDivisor(num1, num2)

	fmt.Printf("Greatest Common Divisor of %d and %d is : %d", num1, num2, res)

}
