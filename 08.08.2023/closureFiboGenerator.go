/*
you want to create a fibonacci number generator, but you don’t want anyone else to have access to that data
(so they can’t accidentally change it). You can use closures to achieve this.
*/

package main

import "fmt"

func makeFiboGenrator() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}

func main() {
	fiboGenerator := makeFiboGenrator()

	for i := 0; i < 10; i++ {
		fmt.Println(fiboGenerator())
	}
}
