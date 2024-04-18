package main

import "fmt"

func main() {
	// Print a greeting message
	fmt.Println("Hello Golang")

	// Loop from 0 to 10 and print each number
	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	// Declare and initialize a slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Loop through the slice and print the index and value of each element
	for index := 0; index < len(numbers); index++ {
		fmt.Printf("The Value of %d th index is %v\n", index, numbers[index])
	}

	// Print a message to indicate the start of the range loop
	fmt.Println("Range Loop")

	// Use a range loop to iterate over the slice and print the index and value of each element
	for index, value := range numbers {
		fmt.Printf("The Value of %d th index is %v\n", index, value)
	}
}
