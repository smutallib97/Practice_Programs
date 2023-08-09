package main

import (
	"fmt"
	"io"
	"os"
)

func createFile(fileName, content string) {
	file, err := os.Create("Content.txt")
	if err != nil {
		panic("file not found")
	}
	defer file.Close()

	_, err := file.WriteString(content)
	if err != nil {
		fmt.Printf("Error ")
	}
}

func readFile() {
	fileName, err := io.ReadAll("Content.txt")
	if err != nil {
		panic("Error while opening file for reading")
	}
	defer fileName.Close()
}
func writeFile() {

}
func main() {

}
