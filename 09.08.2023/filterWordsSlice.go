package main

import (
	"fmt"
	"strings"
)

func filterWords(wordsSlice []string, filter func(string) bool) []string {
	var filteredSliceOfWords []string
	for _, words := range wordsSlice {
		if filter(words) {
			filteredSliceOfWords = append(filteredSliceOfWords, words)
		}
	}
	return filteredSliceOfWords
}
func main() {
	sliceOfWords := []string{"Apple", "Ball", "Cat", "Animal", "Carrot", "Antiseptic"}

	fmt.Println("Original Slice of Words:", sliceOfWords)

	w := "A"

	filteredWords := func(s string) bool {
		return strings.HasPrefix(s, w)
	}

	filteredWordsSlice := filterWords(sliceOfWords, filteredWords)

	fmt.Println("A character filtered slice is: ", filteredWordsSlice)

}
