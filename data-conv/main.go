package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Print a greeting
	fmt.Println("Hello golang")

	// Declare and initialize an integer variable
	var num1 int = 10

	// Print the value of num1
	fmt.Println("Number is:", num1)

	// Print the type of num1 using Printf and %T verb
	fmt.Printf("Type of num1 is: %T \n", num1)

	// Convert num1 to float64 and assign it to num2
	var num2 float64 = float64(num1)

	// Print the type of num2
	fmt.Printf("The Type of num2 is: %T \n", num2)

	// Declare and initialize an integer variable
	num3 := 123

	// Convert num3 to string using strconv.Itoa
	str := strconv.Itoa(num3)

	// Print the string representation of num3
	fmt.Println("The String is:", str)

	// Print the type of str
	fmt.Printf("The Type of str is: %T \n", str)

	// Convert string to integer using strconv.Atoi
	strInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		// Print the converted integer
		fmt.Println("Converted integer from string is:", strInt)
	}

	// Convert string to float64 using strconv.ParseFloat
	strFloat := "3.14"
	floatVal, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
	} else {
		// Print the converted float64
		fmt.Printf("Converted float from string is: %f\n", floatVal)
	}
}
