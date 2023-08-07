package main

import "fmt"

type Student struct {
	RollNo         int
	Name           string
	PhysicsMarks   float32
	ChemistryMarks float32
	MathMarks      float32
	TotalMarks     float32
	Percentage     float32
}

//lculateTotalMarks
func calculateTotalMarks(students []Student) {

	for i := range students {
		totalMarks := students[i].PhysicsMarks + students[i].ChemistryMarks + students[i].MathMarks
		students[i].TotalMarks = totalMarks
	}
}

//CalculatePercentage
func calculatePercentage(students []Student) {

	for i := range students {
		percentage := (students[i].TotalMarks / 300.0) * 100.0
		students[i].Percentage = percentage
	}
}

func calculateThresholdStudents(students []Student, threshold float32) []Student {
	var aboveThresholdStudents []Student
	for _, student := range students {
		if student.Percentage >= threshold {
			aboveThresholdStudents = append(aboveThresholdStudents, student)
		}
	}
	return aboveThresholdStudents
}

// func findHighestPercentageFromStudents(students []Student)  {

// 	for _, student := range students {
// 		if student.Percentage >= highestPercentage {

// 		}
// 	}
// }

func main() {

	studentDetails := []Student{
		{RollNo: 1, Name: "Salim", PhysicsMarks: 75, ChemistryMarks: 85, MathMarks: 75},
		{RollNo: 2, Name: "Ramesh", PhysicsMarks: 70, ChemistryMarks: 80, MathMarks: 75},
		{RollNo: 3, Name: "Mahesh", PhysicsMarks: 72, ChemistryMarks: 78, MathMarks: 80},
		{RollNo: 4, Name: "Akhilesh", PhysicsMarks: 75, ChemistryMarks: 80, MathMarks: 70},
		{RollNo: 5, Name: "Akram", PhysicsMarks: 54, ChemistryMarks: 60, MathMarks: 70},
		{RollNo: 6, Name: "Shahid", PhysicsMarks: 58, ChemistryMarks: 72, MathMarks: 68},
	}
	calculateTotalMarks(studentDetails)
	calculatePercentage(studentDetails)
	fmt.Println("Student Details:")
	for _, student := range studentDetails {
		fmt.Printf("Roll No: %d\n", student.RollNo)
		fmt.Printf("Student Name: %s\n", student.Name)
		fmt.Printf("Physics Marks:%.2f\n", student.PhysicsMarks)
		fmt.Printf("Chemistry Marks: %.2f\n", student.ChemistryMarks)
		fmt.Printf("Math Marks: %.2f\n", student.MathMarks)
		fmt.Printf("Totalmarks : %.2f\n", student.TotalMarks)
		fmt.Printf("Percentage : %.2f\n\n\n", student.Percentage)

	}
	var value float32
	fmt.Print("Enter threshold: ")
	fmt.Scan(&value)

	thresholdValue := calculateThresholdStudents(studentDetails, value)

	fmt.Printf("Number of students above or equal to %.2f : %d", value, len(thresholdValue))

}
