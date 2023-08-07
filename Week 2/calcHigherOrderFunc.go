/*Design a higher-order function called calculator that takes an operator as a string
("add", "subtract", "multiply", "divide") and returns a corresponding function that performs
the specified operation on two integer operands. Test the calculator function with various operations.*/

package main

import "fmt"

func calculator(operator string) func(int, int) int {
	switch operator {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "subtract":
		return func(a, b int) int {
			return a - b
		}
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	case "divide":
		return func(a, b int) int {
			if b == 0 {
				panic("division by zero")
			}
			return a / b
		}
	default:
		panic("unsupported operator")
	}
}

func main() {
	addFunc := calculator("add")
	subtractFunc := calculator("subtract")
	multiplyFunc := calculator("multiply")
	divideFunc := calculator("divide")

	a, b := 10, 5

	fmt.Printf("%d + %d = %d\n", a, b, addFunc(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, subtractFunc(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiplyFunc(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, divideFunc(a, b))
}
