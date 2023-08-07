/*Implement a switch statement in Go to print different messages based on the day of the week.*/

package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()

	switch day {
	case time.Monday:
		fmt.Println("It's Monday.")
	case time.Tuesday:
		fmt.Println("It's Tuesday.")
	case time.Wednesday:
		fmt.Println("It's Wednesday.")
	case time.Thursday:
		fmt.Println("It's Thursday.")
	case time.Friday:
		fmt.Println("It's Friday.")
	case time.Saturday:
		fmt.Println("It's Saturday.")
	case time.Sunday:
		fmt.Println("It's Sunday.")
	default:
		fmt.Println("Invalid day.")
	}
}
