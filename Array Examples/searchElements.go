package main

import "fmt"

func main() {
	arrayOfElements := []int{1, 2, 3, 4, 5, 6, 7}
	var searchElement int
	fmt.Print("Enter a number for searching in array:")
	fmt.Scan(&searchElement)

	for i := 0; i < len(arrayOfElements); i++ {
		if arrayOfElements[i] == searchElement {
			fmt.Printf("\n%d element found at index no. : %d", searchElement, i)
		}
	}

}
