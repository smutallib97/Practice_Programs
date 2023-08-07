// Implement a function to check if a number is prime.
package main

import (
	"fmt"
	"math"
)

func isPrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}
	if num%2 == 0 {
		return false
	}
	if num <= 3 {
		return true // 2 and 3 is prime number
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}

}
func main() {
	var num1 int
	fmt.Println("Check wheather a number is Prime or not")
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num1)

	if isPrimeNumber(num1) {
		fmt.Printf("%d is a prime number\n", num1)
	} else {
		fmt.Printf("%d is not a prime number\n", num1)
	}

}
