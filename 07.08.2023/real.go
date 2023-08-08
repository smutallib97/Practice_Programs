package main

import "fmt"

// type Money struct{   //userdefined data type not comparable
// 	a string
// }

func main() {

	//var a Money
	var a string
	fmt.Print("Enter a chice:")
	fmt.Scan(&a)

	switch a {
	case "1":
		fmt.Println("Having 100 rupees")
		break
	case "2":
		fmt.Println("Having 500 rupees")
		break
	default:
		fmt.Println("Doesn't have money")
	}
}
