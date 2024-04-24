package main

import (
	"fmt"
)

// add function takes two integers, a and b, and returns their sum.
func add(a int, b int) int {
	return a + b
}

func main() {
	// Print a message indicating the start of the program.
	fmt.Println("Starting of the Program")

	// Call the add function with arguments 5 and 6 and store the result in data.
	data := add(5, 6)

	// Defer printing the sum stored in data until the end of the main function.
	defer fmt.Println("The Sum is:", data)

	// Defer printing a message indicating the middle of the program until the end of the main function.
	defer fmt.Println("Middle of the program")

	// Print a message indicating the end of the program.
	fmt.Println("End of the program")
}
