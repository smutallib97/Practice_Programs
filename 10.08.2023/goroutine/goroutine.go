package main

import "fmt"

func hello(name string) {
	fmt.Printf("Hello %s", name)
}
func main() {

	go hello("Mayur")
	go hello("Tauqeer")
	go hello("Sadique")
	go hello("Sir")
}
