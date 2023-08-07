//Given a string, count the number of words in it.

package main

import (
	"fmt"
	"strings"
)

func countWords(input string) int {
	words := strings.Fields(input)

	return len(words)
}

func main() {
	string1 := "welcome to Bridgelabz LLP"

	wordCount := countWords(string1)
	fmt.Printf("Number of words in the string: %d\n", wordCount)
}
