package main

import "fmt"

func filterMap(data []int, filter func(int) int) []int {
	mapped := make([]int, len(data))

	for i, e := range data {
		mapped[i] = filter(e)
	}
	return mapped
}

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := filterMap(vals, func(i int) int {
		return i * i
	})

	fmt.Println(res)

}
