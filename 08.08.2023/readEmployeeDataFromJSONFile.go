/*
Write a Go program to read employees data from JSON file.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Employee struct {
	Emp_Id   int
	Emp_Name string
	Emp_DOJ  string
}

func main() {

	fileName, err := os.Open("data2.json")
	if err != nil {
		log.Fatal(err)
	}

	defer fileName.Close()

	data, err := io.ReadAll(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var result []Employee

	jsonErr := json.Unmarshal(data, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(result)
}
