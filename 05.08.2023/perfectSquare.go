//Create a function to check if a number is a perfect square.

package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	root := int(math.Sqrt(float64(num)))
	return root*root == num
}

func main() {
	number := 25
	if isPerfectSquare(number) {
		fmt.Printf("%d is a perfect square.\n", number)
	} else {
		fmt.Printf("%d is not a perfect square.\n", number)
	}
}
