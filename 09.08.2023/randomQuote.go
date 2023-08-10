/*
Write a program that generates a random quote of the day from a predefined list.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quotes := []string{
		"Life is 10% what happens to us and 90% how we react to it. - Charles R. Swindoll",
		"The only way to do great work is to love what you do. - Steve Jobs",
		"In three words I can sum up everything I've learned about life: it goes on. - Robert Frost",
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
		"Don't watch the clock; do what it does. Keep going. - Sam Levenson",
	}

	rand.Seed(time.Now().UnixNano())
	randomQuote := rand.Intn(len(quotes))

	fmt.Println("Quote of the Day: ", randomQuote)
	fmt.Println(quotes[randomQuote])
}
