//Implement a function to reverse an integer.

package main

import (
	"fmt"
	"strconv"
)

func reverseInteger(number int) int {
	numberString := strconv.Itoa(number)
	var reversedString string

	for i := len(numberString) - 1; i >= 0; i-- {
		reversedString += string(numberString[i])
	}
	reversedNumber, _ := strconv.Atoi(reversedString)

	return reversedNumber
}
func main() {
	//orgNumber := 987654321
	var orgNumber int
	fmt.Printf("Enter Number: ")
	fmt.Scanf("%d", &orgNumber)
	reversedInteger := reverseInteger(orgNumber)

	//fmt.Printf("Original Interger Number is: %d\n ", orgNumber)
	fmt.Printf("Reversed Integer Number is: %d\n", reversedInteger)

}
