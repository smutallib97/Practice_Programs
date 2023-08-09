package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName, err := os.Create("filewrite.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileName.Close()

	words := []string{"Magic", "Happens", "When"}

	for _, word := range words {
		_, err := fileName.WriteString(word + " ")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Done")
}
