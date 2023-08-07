//Write a function to calculate the power of a number.

package main

import "fmt"

func powerOfNumber(base, exponent int) int {
	res := 1
	for i := 0; i < exponent; i++ {
		res *= base
	}
	return res
}
func main() {
	base, exponent := 2, 3

	result := powerOfNumber(base, exponent)
	fmt.Printf("%d raised to the power of %d is: %d\n", base, exponent, result)

}
