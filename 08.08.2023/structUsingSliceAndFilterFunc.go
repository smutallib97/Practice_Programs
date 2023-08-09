package main

import "fmt"

type User struct {
	Name       string
	Occupation string
	City       string
}

func filterSlice(data []User, filter func(User) bool) []User {
	filteredSlice := make([]User, 0)

	for _, user := range data {
		if filter(user) {
			filteredSlice = append(filteredSlice, user)
		}
	}
	return filteredSlice
}

func main() {

	users := []User{
		{Name: "Ramesh", Occupation: "Software Developer", City: "Pune"},
		{Name: "Rajesh", Occupation: "Driver", City: "Nashik"},
		{Name: "Akhilesh", Occupation: "Gardner", City: "Patna"},
		{Name: "Amit", Occupation: "Youtuber", City: "Dehli"},
	}

	city := "Patna"

	userFilter := filterSlice(users, func(u User) bool {
		return u.City == city
	})

	fmt.Println(userFilter)
}
