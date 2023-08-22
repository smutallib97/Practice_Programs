/*package main

import "fmt"

type EmployeeData struct {
	Name       string
	Age        int
	Salary     float64
	Department string
}

func TransFormEmployee(EmployeeOne EmployeeData, transformations ...func(EmployeeData) EmployeeData) EmployeeData {
	for _, transform := range transformations {
		EmployeeOne = transform(EmployeeOne)

	}
	return EmployeeOne

}

func DoubleSalary(e EmployeeData) EmployeeData {
	e.Salary *= 2
	return e

}

func Promote(e EmployeeData) EmployeeData {
	e.Age++
	e.Salary *= 2
	return e
}

func Transfer(e EmployeeData, newDepartment string) EmployeeData {
	newDepartment = e.Department
	return e

}

func main() {

	EmployeeOne := EmployeeData{
		Name:       "Shephali",
		Age:        22,
		Salary:     500000,
		Department: "Finance",
	}

	/*transformedEmployee := TransFormEmployee(EmployeeOne, DoubleSalary, Promote,func (emp EmployeeData)EmployeeData  {
		return Transfer(EmplEmployeeOne, "HR")
	}
	fmt.Println("Original Employee:", EmployeeOne)
	fmt.Println("Transformed Employee:", transformedEmployee)*/

/*trasformation := []func(EmployeeData) EmployeeData{
	DoubleSalary,
	Promote,
	func(emp EmployeeData) EmployeeData {
		return Transfer(emp, "Testing")
	},
}
transformedEmployee := TransFormEmployee(EmployeeOne, trasformation)*/

/*trasformation := []func(EmployeeData) EmployeeData{
		DoubleSalary,
		Promote,
		func(emp EmployeeData) EmployeeData {
			return Transfer(emp, "Testing")
		},
	}

	transformedEmployee := TransFormEmployee(EmployeeOne, trasformation...)
	fmt.Println("Original Employee:", EmployeeOne)
	fmt.Println("Transformed Employee:", transformedEmployee)
}*/

package main

import (
	"fmt"
)

type EmployeeData struct {
	Name       string
	Age        int
	Salary     float64
	Department string
}

func TransFormEmployee(EmployeeOne EmployeeData, transformations ...func(EmployeeData) EmployeeData) EmployeeData {
	for _, transform := range transformations {
		EmployeeOne = transform(EmployeeOne)
	}
	return EmployeeOne
}

func DoubleSalary(e EmployeeData) EmployeeData {
	e.Salary *= 2
	return e
}

func Promote(e EmployeeData) EmployeeData {
	e.Age++
	e.Salary *= 2
	return e
}

func Transfer(e EmployeeData, newDepartment string) EmployeeData {
	e.Department = newDepartment
	return e
}

func main() {
	EmployeeOne := EmployeeData{
		Name:       "Shephali",
		Age:        22,
		Salary:     500000,
		Department: "Finance",
	}

	trasformation := []func(EmployeeData) EmployeeData{
		DoubleSalary,
		Promote,
		func(emp EmployeeData) EmployeeData {
			return Transfer(emp, "Testing")
		},
	}

	transformedEmployee := TransFormEmployee(EmployeeOne, trasformation...)
	fmt.Println("Original Employee:", EmployeeOne)
	fmt.Println("Transformed Employee:", transformedEmployee)
}
