package main

import (
	"fmt"
)

// Person struct represents a person with first name, last name, and age.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Printing a message to indicate that we're learning about structs.
	fmt.Println("Learning Structs...")

	// Declaring a variable of type Person
	var person1 Person

	// Printing the initial value of person1 (it will be zero-valued)
	fmt.Println("Person1:", person1)

	// Assigning values to the fields of person1
	person1.FirstName = "Swapnil"
	person1.LastName = "Khatik"
	person1.Age = 22

	// Printing person1 after assigning values
	fmt.Println("Person1:", person1)

	// Declaring and initializing person2 using struct literal
	person2 := Person{
		FirstName: "Ritesh",
		LastName:  "Khore",
		Age:       22,
	}

	// Printing person2
	fmt.Println("Person2:", person2)

	// Printing the type of person2
	fmt.Printf("The type of person2 is %T\n", person2)

	// Declaring and initializing person3 using the new function
	person3 := &Person{
		FirstName: "Pratik",
		LastName:  "Ghadge",
		Age:       22,
	}

	// Printing person3
	fmt.Println("Person3:", *person3)

	// It's more idiomatic to use the zero-value initialization for pointers to structs.
	// If you want to initialize it with zero values, you can simply do:
	// person3 := &Person{}

	// You can access the fields of a pointer to a struct directly.
	// So instead of person3.FirstName, you can directly do person3.FirstName.
	// This is more concise and clear in most cases.
	// e.g., person3.FirstName = "Pratik"

	// If you want to keep the initialization using new() function, you can use it like this:
	// person3 := new(Person)
	// Then you can directly assign values to its fields: person3.FirstName = "Pratik", etc.

}
