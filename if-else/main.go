package main

import "fmt"

func main() {
	// Print a welcome message
	fmt.Println("Hello Golang")

	// Prompt the user to enter a value for x
	var x int
	fmt.Print("Enter Value of x: ")
	fmt.Scan(&x)

	// Check if the number is even
	if x%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	// Additional conditions
	// Check if the number is positive, negative, or zero
	if x > 0 {
		fmt.Println("Positive Number")
	} else if x < 0 {
		fmt.Println("Negative Number")
	} else {
		fmt.Println("Zero")
	}

	// Check if the number is a multiple of 3
	if x%3 == 0 {
		fmt.Println("Multiple of 3")
	} else {
		fmt.Println("Not a Multiple of 3")
	}

	// Examples of combining conditions
	if x > 0 && x%2 == 0 {
		fmt.Println("Positive Even Number")
	}

	if x < 0 || x%2 != 0 {
		fmt.Println("Negative or Odd Number")
	}

}
