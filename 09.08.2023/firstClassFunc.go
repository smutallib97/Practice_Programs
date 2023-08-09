/*First Class Functions:
->Functions can be passed in as arguments to other function
->Functions can be returned from other functions
->Functions can be stored in data structures

=> Function will be treated as values
*/

package main

import "fmt"

func firstClassFunc() {
	addFunc := func(a, b int) int {
		return a + b
	}
	substractionFunc := func(a, b int) int {
		return a - b
	}
	functionsArray := []func(int, int) int{addFunc, substractionFunc}

	for _, function := range functionsArray {
		fmt.Println(function(10, 20))
	}
}
func main() {
	fmt.Println("Addtion and Substraction of two numbers:")
	firstClassFunc()
}
