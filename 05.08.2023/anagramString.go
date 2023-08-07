//Implement a function to check if a string is an anagram of another string.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func isAnagram(str1, str2 string) bool {
	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	sortedStr1 := sortString(str1)
	sortedStr2 := sortString(str2)

	return sortedStr1 == sortedStr2
}

func sortString(s string) string {
	charSlice := strings.Split(s, "")
	sort.Strings(charSlice)
	return strings.Join(charSlice, "")
}

func main() {
	word1 := "Silent"
	word2 := "Listen"

	if isAnagram(word1, word2) {
		fmt.Printf("%s and %s are anagrams.\n", word1, word2)
	} else {
		fmt.Printf("%s and %s are not anagrams.\n", word1, word2)
	}
}
