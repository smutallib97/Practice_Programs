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

// func calculateThresholdStudents(students []Student, threshold float32) []Student {
// 	var aboveThresholdStudents []Student
// 	for _, student := range students {
// 		if student.Percentage >= threshold {
// 			aboveThresholdStudents = append(aboveThresholdStudents, student)
// 		}
// 	}
// 	return aboveThresholdStudents
// }

// func findHighestPercentageFromStudents(students []Student)  {

// 	for _, student := range students {
// 		if student.Percentage >= highestPercentage {

// 		}
// 	}
// }

func filterStudents(students []Student, filterfunc func(Student) bool) []Student {
	var filteredStudents []Student
	for _, student := range students {
		if filterfunc(student) {
			filteredStudents = append(filteredStudents, student)
		}
	}
	return filteredStudents
}

func findTopperStudent(students []Student, compareFunc func(Student, Student) bool) Student {
	topperStudent := students[0]
	for _, student := range students {
		if compareFunc(student, topperStudent) {
			topperStudent = student
		}
	}
	return topperStudent
}

func main() {

	studentDetails := []Student{
		{RollNo: 1, Name: "Salim", PhysicsMarks: 75, ChemistryMarks: 85, MathMarks: 75},
		{RollNo: 2, Name: "Ramesh", PhysicsMarks: 70, ChemistryMarks: 80, MathMarks: 75},
		{RollNo: 3, Name: "Mahesh", PhysicsMarks: 72, ChemistryMarks: 78, MathMarks: 80},
		{RollNo: 4, Name: "Akhilesh", PhysicsMarks: 75, ChemistryMarks: 80, MathMarks: 70},
		{RollNo: 5, Name: "Akram", PhysicsMarks: 54, ChemistryMarks: 60, MathMarks: 70},
		{RollNo: 6, Name: "Shahid", PhysicsMarks: 58, ChemistryMarks: 72, MathMarks: 68},
		{RollNo: 7, Name: "Akram", PhysicsMarks: 40, ChemistryMarks: 40, MathMarks: 40},
		{RollNo: 8, Name: "Sayali", PhysicsMarks: 35, ChemistryMarks: 38, MathMarks: 39},
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
	//filtering students who got lower than 50 percentage
	lowerThan50PercentageFilter := func(student Student) bool {
		return student.Percentage < 50
	}
	lowerThan50PercentageStudents := filterStudents(studentDetails, lowerThan50PercentageFilter)

	fmt.Println("Number of student who got lower than 50 percentage are: ", len(lowerThan50PercentageStudents))

	topperComparator := func(student1, student2 Student) bool {
		return student1.Percentage > student2.Percentage
	}

	topperStudent := findTopperStudent(studentDetails, topperComparator)

	fmt.Println("\n\n Topper Student From Class is: ")
	fmt.Printf("Roll No: %d\n", topperStudent.RollNo)
	fmt.Printf("Student Name: %s\n", topperStudent.Name)
	fmt.Printf("Physics Marks:%.2f\n", topperStudent.PhysicsMarks)
	fmt.Printf("Chemistry Marks: %.2f\n", topperStudent.ChemistryMarks)
	fmt.Printf("Math Marks: %.2f\n", topperStudent.MathMarks)
	fmt.Printf("Totalmarks : %.2f\n", topperStudent.TotalMarks)
	fmt.Printf("Percentage : %.2f\n\n\n", topperStudent.Percentage)

	failedComparator := func(student Student) bool {
		return student.Percentage < 40
	}
	failedStudent := filterStudents(studentDetails, failedComparator)

	fmt.Println("Number of student who got failed :", len(failedStudent))

	physicsMarksComparator := func(student Student) bool {
		return student.PhysicsMarks == 70
	}
	physicsStudent := filterStudents(studentDetails, physicsMarksComparator)

	fmt.Println("Number of students who got 70 marks in physics are: ", len(physicsStudent))

	// var value float32
	// fmt.Print("Enter threshold: ")
	// fmt.Scan(&value)

	// thresholdValue := calculateThresholdStudents(studentDetails, value)

	// fmt.Printf("Number of students above or equal to %.2f : %d", value, len(thresholdValue))

}
