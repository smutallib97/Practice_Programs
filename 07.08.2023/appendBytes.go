package main

import "fmt"

func main() {

	bytes := append([]byte("old"), "falcon"...)

	fmt.Println(bytes)
	fmt.Println(string(bytes))
	//cutting slice
	slice1 := bytes[:2]
	slice2 := bytes[2:5]
	//removing elements from slice
	bytes = append(bytes[:6], bytes[7:8]...)

	fmt.Println("Slice1 is :", slice1)
	fmt.Println("Slice2 is :", slice2)
}
