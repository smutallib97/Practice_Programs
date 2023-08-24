package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type EmployeeData struct {
	EmpID   int    `json:"EmpId"`
	EmpName string `json:"EmpName"`
}

var employeeData = []EmployeeData{}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(employeeData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
func postData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newEmployee EmployeeData
		err := json.NewDecoder(r.Body).Decode(&newEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newEmployee.EmpID = len(employeeData) + 1
		employeeData = append(employeeData, newEmployee)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Employee added successfully")
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func putData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var updatedEmployee EmployeeData
		err := json.NewDecoder(r.Body).Decode(&updatedEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, emp := range employeeData {
			if emp.EmpID == updatedEmployee.EmpID {
				employeeData[i] = updatedEmployee
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Employee updated successfully")
				return
			}
		}
		http.Error(w, "Employee not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func patchData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		var updatedEmployee EmployeeData
		err := json.NewDecoder(r.Body).Decode(&updatedEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, emp := range employeeData {
			if emp.EmpID == updatedEmployee.EmpID {
				if updatedEmployee.EmpName != "" {
					employeeData[i].EmpName = updatedEmployee.EmpName
				}
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Employee updated successfully")
				return
			}
		}
		http.Error(w, "Employee not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func deleteData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		empIDStr := r.URL.Query().Get("EmpId")
		empID, err := strconv.Atoi(empIDStr)
		if err != nil {
			http.Error(w, "Invalid EmpId parameter", http.StatusBadRequest)
			return
		}
		for i, emp := range employeeData {
			if emp.EmpID == empID {
				employeeData = append(employeeData[:i], employeeData[i+1:]...)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "Employee deleted successfully")
				return
			}
		}
		http.Error(w, "Employee not found", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func main() {
	http.HandleFunc("/employeeDetails", getData)
	http.HandleFunc("/newEmployee", postData)
	http.HandleFunc("/updateEmployee", putData)
	http.HandleFunc("/patchEmployee", patchData)
	http.HandleFunc("/deleteEmployee", deleteData)
	log.Println("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
