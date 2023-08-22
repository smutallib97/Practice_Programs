/*
 use higher-order functions and closure functions in Go to create a logical program
 to generate double and triple of a number.
*/

package main

import "fmt"

func multiplyBy(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

func main() {

	var num int
	fmt.Print("Enter a number that you want to double and triple of that value: ")
	fmt.Scan(&num)

	double := multiplyBy(2)
	triple := multiplyBy(3)

	fmt.Printf("Double of %d is: %d \n", num, double(num))
	fmt.Printf("Triple of %d is: %d \n", num, triple(num))
}
