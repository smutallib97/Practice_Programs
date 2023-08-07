//FizzBuzz: Print numbers from 1 to n, but for multiples of 3,
//print "Fizz," for multiples of 5, print "Buzz," and for multiples
//of both 3 and 5, print "FizzBuzz."

package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter a Number")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}
