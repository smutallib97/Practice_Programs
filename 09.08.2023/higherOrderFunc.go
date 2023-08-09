/*Higher Order Function

a function can takes one or more functions as arguments and returns function as its result itself.
*/

package main

import "fmt"

func calculate(a, b int, funcApply func(c, d int) int) int {
	return funcApply(a, b)
}

func higherOrderFunc() {
	multiplication := func(a, b int) int {
		return a * b
	}
	fmt.Println("Product of 10 and 20: ", calculate(10, 20, multiplication))
}

func main() {
	higherOrderFunc()

}
