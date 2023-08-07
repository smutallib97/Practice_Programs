/*
 Write a Go program that opens a file, reads its content, and defers the file closing
 process until the end of the program using an anonymous function. Handle any potential
 errors that may occur during the file access and read operations.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "example.txt"

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	fmt.Printf("File content:\n%s", file)
}
