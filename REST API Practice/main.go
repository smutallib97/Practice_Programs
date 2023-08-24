package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	EmpID   int    `json:"id"`
	EmpName string `json:"empname"`
}

var employeeData = []Employee{
	{EmpID: 101, EmpName: "Sayyed"},
	{EmpID: 102, EmpName: "Tauqeer"},
}

func getData(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, employeeData)
}

func postData(context *gin.Context) {
	var newEmployee Employee
	if err := context.BindJSON(&newEmployee); err != nil {
		return
	}
	employeeData = append(employeeData, newEmployee)
	context.IndentedJSON(http.StatusCreated, employeeData)
}

func main() {
	router := gin.Default()
	fmt.Println("Listening from Main Server")
	router.GET("/employeeData", getData)
	router.POST("/employeeData/newEmployee", postData)
	router.Run("localhost:8080")

}
