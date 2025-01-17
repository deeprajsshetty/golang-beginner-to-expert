package main

import "fmt"

type Student struct {
	Name string
	ID int
	age float64
}

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(studentIDs []int, studentID int) []int {
	return append(studentIDs, studentID)
}

func addStudentStruct(studentStructs []Student, studentStruct Student) []Student {
	return append(studentStructs, studentStruct)
}

func main()  {
	students := []string{}	// Empty String Slice
	result := addStudent(students, "First Student")
	result = addStudent(result, "Second Student")
	result = addStudent(result, "Third Student")
	fmt.Println(result)

	studentIDs := []int{}	// Empty Int Slice
	resultID := addStudentID(studentIDs, 2000)
	resultID = addStudentID(resultID, 300)
	resultID = addStudentID(resultID, 400)
	fmt.Println(resultID)
}

/*
Output:
[First Student Second Student Third Student]
[2000 300 400]
*/

