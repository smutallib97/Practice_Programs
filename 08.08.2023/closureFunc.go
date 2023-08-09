//A closure is a special type of anonymous function that references variables declared outside of the function it self.

package main

import "fmt"

func main() {
	n := 0
	counter := func() int {
		n += 1
		return n
	}
	fmt.Println(counter())
	fmt.Println(counter())
}
