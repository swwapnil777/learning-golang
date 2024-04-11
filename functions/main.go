package main

import (
	"fmt"
	"math"
	"strconv"
)

// sumIntegers takes two integers a and b as input and returns their sum.
func sumIntegers(a int, b int) int {
	return a + b
}

// subtractIntegers takes two integers a and b as input and returns their difference.
func subtractIntegers(a int, b int) int {
	return a - b
}

// multiplyFloats takes two float64 numbers a and b as input and returns their product.
func multiplyFloats(a float64, b float64) float64 {
	return a * b
}

// divideFloats takes two float64 numbers a and b as input and returns their division result.
// It also returns an error if the second number is zero.
func divideFloats(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// power calculates the power of a given number to the specified exponent.
func power(base float64, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// isEven checks if a given integer is even.
func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	var input string
	var err error

	// Get the first number from the user
	fmt.Print("Enter the first integer number (x): ")
	fmt.Scanln(&input)
	x, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the second number from the user
	fmt.Print("Enter the second integer number (y): ")
	fmt.Scanln(&input)
	y, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the float number from the user
	fmt.Print("Enter the float number (z): ")
	fmt.Scanln(&input)
	z, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calculate the sum of x and y
	fmt.Println("The sum of", x, "and", y, "is:", sumIntegers(x, y))

	// Calculate the difference of x and y
	fmt.Println("The difference of", x, "and", y, "is:", subtractIntegers(x, y))

	// Calculate the product of x and z (float numbers)
	fmt.Println("The product of", x, "and", z, "is:", multiplyFloats(float64(x), z))

	// Calculate the division result of x and z (float numbers)
	result, err := divideFloats(float64(x), z)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The division result of", x, "and", z, "is:", result)
	}

	// Calculate the power of x to the y
	fmt.Println(x, "raised to the power of", y, "is:", power(float64(x), float64(y)))

	// Check if x is even
	fmt.Println(x, "is even:", isEven(x))
}
