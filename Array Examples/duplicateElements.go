package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 5, 4, 3, 3}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i] && i != j {
				fmt.Println(arr[i])
			}
		}
	}
}
