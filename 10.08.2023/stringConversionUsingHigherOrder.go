package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertString(numbers []int, conversionFunc func(int) string) []string {
	convertedString := make([]string, len(numbers))

	for i, num := range numbers {
		convertedString[i] = conversionFunc(num)
	}
	return convertedString
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	toString := func(i int) string {
		return strconv.Itoa(i)
	}
	convertedNumbersToString := convertString(numbers, toString)

	fmt.Println("Converted strings:", convertedNumbersToString)
	fmt.Println("Joined strings:", strings.Join(convertedNumbersToString, ", "))
}
