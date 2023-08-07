/*Develop a Temperature Unit Converter Program. It can
convert temperature from centigrade to Fahrenheit or
vice-versa based on user input.
It takes user input to select between 0 and 1. If input is
0, it converts Celsius to Fahrenheit and 1 for F ->C
Then It asks user to enter temperature
It converts to other and displays the same*/

package main

import "fmt"

func celsiusToFehrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32 // coverting into fehrenheit
}

func fehrenheitToCelsius(fehrenheit float64) float64 {
	return (fehrenheit - 32) * 5 / 9
}

func main() {
	var tempCelsius float64
	var tempFehrenheit float64

	fmt.Println("Temprerature Unit Converter")
	fmt.Println("Press 1 to convert Celsius to Temprerature")
	fmt.Println("Press 2 to convert Temprerature to Celsius")

	var choice int
	fmt.Print("Enter your choice 1 or 2: ")
	fmt.Scan(&choice)

	switch choice {

	case 1:

		fmt.Print("Enter temprerature in celsius: ")
		fmt.Scan(&tempCelsius)

		tempFehrenheit = celsiusToFehrenheit(tempCelsius)
		fmt.Printf("%.2f °C is equal to %.2f °F\n", tempCelsius, tempFehrenheit)

	case 2:
		fmt.Print("Enter temprerature in fehrenheit: ")
		fmt.Scan(&tempFehrenheit)

		tempCelsius = fehrenheitToCelsius(tempFehrenheit)
		fmt.Printf("%.2f °F is equal to %.2f °C\n", tempFehrenheit, tempCelsius)

	default:
		fmt.Println("Invalid Input")

	}
}
