package main

import (
	"log"
	"os"
)

func main() {
	srcFile := "data2.json"
	destFile := "data3.json"

	bytesRead, err := os.ReadFile(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(destFile, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}

	println("Done")
}
