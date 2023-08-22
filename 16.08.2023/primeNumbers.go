/*Create a Go program that prints all prime numbers between 1 and n, where n is provided by the user.*/

package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n int
	fmt.Print("Enter a number :")
	fmt.Scan(&n)

	fmt.Printf("Prime numbers between 1 to %d: \n", n)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
