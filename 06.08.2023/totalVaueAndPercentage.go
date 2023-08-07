//Write a Go program that takes two float64 inputs: a total value and a percentage.
//The program should calculate and print the result of applying the percentage to the total value.

package main

import "fmt"

func applyingPercentage(total float64, percentage float64) float64 {
	return total * (percentage / 100)
}

func main() {

	var total_value, percentage float64

	fmt.Print("Enter total value and percentage: ")
	fmt.Scan(&total_value, &percentage)

	result := applyingPercentage(total_value, percentage)

	fmt.Printf("%.2f of %.2f is : %.2f", total_value, percentage, result)

}
