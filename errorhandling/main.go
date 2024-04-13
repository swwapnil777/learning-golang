package main

import (
	"errors"
	"fmt"
)

// divideFunction divides two float32 numbers and returns the result.
// It also checks for division by zero and returns an error if encountered.
func divide(dividend float32, divisor float32) (float32, error) {
	if divisor == 0 {
		// Return an error if the divisor is zero
		return 0, errors.New("denominator must not be zero")
	}

	// Perform the division and return the result
	return dividend / divisor, nil
}

func main() {
	fmt.Println("Welcome to the Error Handling Example")

	// Call the divide function with sample values
	result, err := divide(10, 0)

	if err != nil {
		// Handle the error if division by zero occurs
		fmt.Println("Error:", err)
		return
	}

	// Print the result if no error occurred
	fmt.Println("The division result is:", result)
}
