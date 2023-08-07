//Create a function to check if a slice contains a specific element.

package main

import "fmt"

func isSliceContainSpecificElement(slice []int, element int) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice :", slice1)
	var element int
	fmt.Printf("Enter a element to check element present or not in above slice: ")
	fmt.Scanf("%d", &element)

	if isSliceContainSpecificElement(slice1, element) {
		fmt.Printf("The element %d is present in the slice. \n ", element)
	} else {
		fmt.Printf("The element %d is not present in the slcie. \n", element)
	}

}
