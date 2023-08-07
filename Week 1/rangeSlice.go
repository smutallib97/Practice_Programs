/*Write a Go program that uses the range keyword to iterate over a slice and print its elements.*/

package main

import (
	"fmt"
)

func main() {
	fruits := []string{"apple", "banana", "orange", "grape"}

	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}
