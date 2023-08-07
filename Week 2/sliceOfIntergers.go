/*Create a function in Go that takes a slice of integers as input and modifies
the original slice by doubling the value of each element. Explain how Go handles
passing the slice as an argument and whether the changes to the slice are reflected outside the function.*/

package main

import (
	"fmt"
)

func doubleElements(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	doubleElements(numbers)
	fmt.Println("Modified slice:", numbers)
}
