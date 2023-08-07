/*Write a Go function called counter that returns another function.
The returned function should be a closure that keeps track of how many times
it has been called and prints the count. Test the counter function by calling
 it multiple times and invoking the returned function.*/

package main

import (
	"fmt"
)

func counter() func() {
	count := 0
	return func() {
		count++
		fmt.Printf("Function called %d times.\n", count)
	}
}

func main() {
	fn := counter()

	fn()
	fn()
	fn()
}
