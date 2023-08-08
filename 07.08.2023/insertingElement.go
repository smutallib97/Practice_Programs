package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5}

	i := 2
	e := 3

	a = append(a[:i], append([]int{e}, a[i:]...)...)

	fmt.Println(a)
}
