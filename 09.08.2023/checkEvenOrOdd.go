/*
Write a program to determine if a number is even or odd.
*/

package main

import "fmt"

func checkForEvenOdd(number int) {
	if number%2 == 0 {
		fmt.Printf("%d is a even number ", number)
	} else {
		fmt.Printf("%d is a Odd number", number)
	}
}

func main() {

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	checkForEvenOdd(num)

}
