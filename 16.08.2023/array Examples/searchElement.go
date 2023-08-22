package main

import "fmt"

func main() {
	arr := []int{1, 5, 6, 8, 2, 4, 3}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] <= arr[j] && arr[i] == arr[j] {
				fmt.Println(arr[i])
			}
		}
	}
}
