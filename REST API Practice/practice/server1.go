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

func main() {
	router := gin.Default()
	fmt.Println("Listening from Server 1")
	router.GET("/employeeData", getData)
	router.Run("localhost:8080")

}
