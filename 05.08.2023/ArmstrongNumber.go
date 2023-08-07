//Write a function to check given number is an Armstrong number.

package main

import (
	"fmt"
	"math"
)

func isArmstrong(num int) bool {
	n := num
	sum := 0
	digits := int(math.Log10(float64(num))) + 1

	for num > 0 {
		digit := num % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		num /= 10
	}

	return sum == n
}

func main() {
	num := 153
	if isArmstrong(num) {
		fmt.Println(num, "is an Armstrong number.")
	} else {
		fmt.Println(num, "is not an Armstrong number.")
	}
}
