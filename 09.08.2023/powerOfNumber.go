package main

import (
	"fmt"
	"math"
)

func powerFunc(number float64) func(float64) float64 {
	return func(x float64) float64 {
		return math.Pow(x, number)
	}
}

func main() {

	square := powerFunc(2)
	cube := powerFunc(3)
	fmt.Println("Square and cube of 2 is:")
	fmt.Println(square(2))
	fmt.Println(cube(2))

}
