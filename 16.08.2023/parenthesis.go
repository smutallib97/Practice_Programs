/*Create a Go function to validate whether a given string containing parentheses is properly balanced.*/

package main

import "fmt"

func isBalanced(str string) bool {
	stack := []rune{}

	for _, char := range str {
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else if char == '}' || char == ']' || char == ')' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if char == ')' && top != '(' || char == ']' && top != '[' || char == '}' && top != '{' {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	validationData := []string{
		"{}",
		"[]",
		"()",
		"}{[()]}",
	}

	for _, data := range validationData {
		if isBalanced(data) {
			fmt.Printf("%s is balanced \n", data)
		} else {
			fmt.Printf("%s is not balanced \n", data)
		}
	}

}
