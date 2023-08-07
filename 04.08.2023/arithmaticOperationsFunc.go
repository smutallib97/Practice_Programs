package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func substraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}
func division(a, b int) int {
	return a / b
}

func main() {

	fmt.Println("Arithmatic Operations :")
	var x, y, z int
	fmt.Printf("Enter two numbers:")
	fmt.Scanf("%d\n %d:", &x, &y)

	z = addition(x, y)
	fmt.Println("Addition of two numbers:", z)
	z = substraction(x, y)
	fmt.Println("Substraction of two numbers:", z)
	z = multiplication(x, y)
	fmt.Println("Multiplication of two numbers:", z)
	z = division(x, y)
	fmt.Println("Division of two numbers:", z)

}
