package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	numbers = append(numbers, 1, 2, 3)

	fmt.Println(numbers)
	fmt.Println("Length of Slice", len(numbers))
	fmt.Println("Capacity of slice", cap(numbers))

}
