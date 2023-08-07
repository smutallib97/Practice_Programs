//Given a string, write a function to determine if it is a palindrome.

package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(word string) {
	if word == reverse(word) {
		fmt.Printf("%s is Palindrome", word)
	} else {
		fmt.Printf("%s is not Palindrome", word)
	}
}

func main() {
	var str string
	fmt.Println("Check Weather a string is Palindrome or not")
	fmt.Printf("Enter a string: ")
	fmt.Scanf("%s", &str)

	isPalindrome(str)
}
