//Create a function to check if a year is a leap year.

package main

import "fmt"

func isLeapYear(year int) {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d is leap year", year)
	} else {
		fmt.Printf("%d is not leap year", year)
	}
}

func main() {
	var year int
	fmt.Printf("Enter Year: ")
	fmt.Scanf("%d", &year)

	isLeapYear(year)

}
