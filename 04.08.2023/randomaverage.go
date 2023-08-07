/*Go program that generates 10 random numbers between 1 to 20 and then calculates their average*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(min, max, count int) []int {
	rand.Seed(time.Now().UnixNano())
	randomNumbers := make([]int, count)
	for i := 0; i < count; i++ {
		randomNumbers[i] = rand.Intn(max-min+1) + min
	}
	return randomNumbers
}
func calculateAverage(numbers []int) float64 {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}
func main() {
	min := 1
	max := 20
	count := 10

	randomNumbers := generateRandomNumbers(min, max, count)
	fmt.Println("Generated Random numbers: ", randomNumbers)

	averageOfRandomNumbers := calculateAverage(randomNumbers)
	fmt.Println("Average of Random Numbers: ", averageOfRandomNumbers)
}
