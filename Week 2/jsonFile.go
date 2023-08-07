/*Write a Go function that reads data from a JSON file named "data.json" and
returns it as a Go data structure. The function should handle any file access
errors and return an appropriate error if the file doesn't exist or cannot be read.*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Data struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func ReadDataFromFile(filename string) (Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Data{}, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	dataBytes, err := io.ReadAll(file)
	if err != nil {
		return Data{}, fmt.Errorf("error reading file: %w", err)
	}

	var data Data
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return Data{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return data, nil
}

func main() {
	data, err := ReadDataFromFile("data.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", data.Name)
	fmt.Println("Age:", data.Age)
	fmt.Println("Email:", data.Email)
}
