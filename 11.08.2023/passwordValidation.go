/*
Create a password validation program that checks whether a given password meets certain criteria using higher-order functions.

    Implement a function called IsStrong that takes in a string (password) and a list of validation functions.
  Each validation function should take a string (password) and return a boolean indicating whether the password
  passes that specific validation rule.

    Create at least three validation functions:
        HasMinimumLength that checks if the password has a minimum length of 8 characters.
        HasUppercase that checks if the password contains at least one uppercase letter.
        HasSpecialCharacter that checks if the password contains at least one special character (e.g., !, @, #, $, etc.).

    The IsStrong function should iterate through the list of validation functions and return true if the password
  passes all the validation rules, and false otherwise.
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isStrong(password string, validation []func(string) bool) bool {
	for _, validate := range validation {
		if !validate(password) {
			return false
		}
	}
	return true
}

func hasMinimumLength(password string) bool {
	return len(password) <= 8
}
func hadUpperCase(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}
func hasSpecialCharacter(password string) bool {
	specialCharacters := "!@#$%^&*()_-+=[]{}|;:,.<>?/"
	for _, char := range password {
		if strings.ContainsRune(specialCharacters, char) {
			return true
		}
	}
	return false
}

func main() {
	validationFuncs := []func(string) bool{hasMinimumLength, hadUpperCase, hasSpecialCharacter}

	password1 := "Strong@2020"
	password2 := "weakpass101"
	password3 := "Weakerpass"
	password4 := "12457897"
	password5 := "A@liya"

	fmt.Printf("Password 1 is strong:%v \n", isStrong(password1, validationFuncs))
	fmt.Printf("Password 2 is strong:%v \n", isStrong(password2, validationFuncs))
	fmt.Printf("Password 3 is strong:%v \n", isStrong(password3, validationFuncs))
	fmt.Printf("Password 4 is strong:%v \n", isStrong(password4, validationFuncs))
	fmt.Printf("Password 5 is strong:%v \n", isStrong(password5, validationFuncs))
}
