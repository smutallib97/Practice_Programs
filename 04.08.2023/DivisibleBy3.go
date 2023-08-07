//Slice of Nos. 7 9 20 18 15 27 find out which Nos Divisible by 3

package main

import "fmt"

func main() {

	Slice := []int{7, 9, 20, 18, 15, 27}

	fmt.Println("Elements Divisible 3")

	for _, num := range Slice {
		if num%3 == 0 {
			fmt.Println(num)
		}
	}
}
