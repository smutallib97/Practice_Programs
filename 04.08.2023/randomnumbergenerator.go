package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() (int, int, int) {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)
	z := rand.Intn(10)

	return x, y, z
}

func main() {
	randomNum1, randomNum2, randomNum3 := randomNumber()

	fmt.Println("Three Random Numbers:", randomNum1, randomNum2, randomNum3)

}
