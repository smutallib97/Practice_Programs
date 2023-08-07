//Write a function to convert Fahrenheit to Celsius.

package main

import "fmt"

func fehrenheitToCelsius(fehrenheit float64) float64 {
	return (fehrenheit - 32) * 5 / 9
}
func main() {
	var tempCelsius float64
	var tempFehrenheit float64
	fmt.Print("Enter temprerature in fehrenheit: ")
	fmt.Scan(&tempFehrenheit)

	tempCelsius = fehrenheitToCelsius(tempFehrenheit)
	fmt.Printf("%.2f °F is equal to %.2f °C\n", tempFehrenheit, tempCelsius)

}
