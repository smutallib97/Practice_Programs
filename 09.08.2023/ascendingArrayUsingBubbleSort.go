package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func main() {
	numbersArray := []int{54, 18, 5, 14, 36, 12, 3}
	fmt.Println("Before sorting the array :", numbersArray)
	bubbleSort(numbersArray)
	fmt.Println("after sorting array in Ascending order array is: ", numbersArray)

}
