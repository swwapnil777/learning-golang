package main

import "fmt"

func main() {
	// Print a greeting message
	fmt.Println("Hello Golang")

	// Create a map to store student grades
	studentGrades := make(map[string]int)

	// Add grades for students
	studentGrades["Swapnil"] = 90
	studentGrades["Ritesh"] = 90
	studentGrades["Pratik"] = 90
	studentGrades["Atul"] = 90

	// Add a new student and their grade
	studentGrades["Kaustub"] = 89

	// Delete a student from the map
	delete(studentGrades, "Kaustub")

	// Retrieve grades for "Swapnil" and check if the student exists
	grades, exists := studentGrades["Swapnil"]

	// Print the grades of Swapnil
	fmt.Println("The Grades of Swapnil is", grades)

	// Print whether Swapnil exists in the map
	fmt.Println("Swapnil Exists:", exists)

	// Iterate over the map and print each key-value pair
	fmt.Println("Student Grades:")
	for name, grade := range studentGrades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	// Check the length of the map
	fmt.Println("Number of students:", len(studentGrades))

	// Check if a specific student exists in the map
	if _, exists := studentGrades["Ritesh"]; exists {
		fmt.Println("Ritesh is in the map.")
	} else {
		fmt.Println("Ritesh is not in the map.")
	}
}
