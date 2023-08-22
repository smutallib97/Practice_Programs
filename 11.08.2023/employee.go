/*
Imagine you have a dataset of employee information, and you need to perform various data transformations on it. Implement a program that demonstrates

	  function composition for these transformations.

	    Define a struct Employee with fields such as Name, Age, Salary, and Department.

	    Create functions to perform the following transformations:
	        doubleSalary that doubles the salary of an employee.
	        promote that increases the age of an employee by 1 and doubles their salary.
	        transfer that changes the department of an employee to the provided department.

	    Implement a function called transformEmployee that takes an Employee and a list of transformation functions.
		It applies each transformation
	    function in sequence to the employee data.
*/
package main

import "fmt"

type Employee struct {
	Name       string
	Age        int
	Salary     float32
	Department string
}

func doubleSalary(emp Employee) Employee {
	emp.Salary *= 2.0
	return emp
}

func promote(emp Employee) Employee {
	emp.Age++
	emp.Salary *= 2.0
	return emp
}

func transfer(emp Employee, newDepartment string) Employee {
	emp.Department = newDepartment
	return emp
}

func transformEmployee(emp Employee, transformations []func(Employee) Employee) Employee {
	for _, temp := range transformations {
		emp = temp(emp)
	}
	return emp
}

func main() {
	employee := Employee{Name: "Ashok", Age: 24, Salary: 450000, Department: "Development"}
	trasformation := []func(Employee) Employee{
		doubleSalary,
		promote,
		func(emp Employee) Employee {
			return transfer(emp, "Testing")
		},
	}
	transformedEmployee := transformEmployee(employee, trasformation)

	fmt.Printf("Original Employee: %+v\n", employee)
	fmt.Printf("After transfer to another department : %+v\n", transformedEmployee)

}
