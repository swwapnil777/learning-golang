package main

import "fmt"

func main() {
	// Initialize a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Print the size (length) of the slice
	fmt.Println("Initial Size:", len(numbers)) // Outputs: Initial Size: 5

	// Append more integers to the slice
	numbers = append(numbers, 6, 7, 8, 9, 10)

	// Print the new size (length) of the slice
	fmt.Println("Updated Size:", len(numbers)) // Outputs: Updated Size: 10

	// Print the content of the slice
	fmt.Println("Numbers:", numbers) // Outputs: Numbers: [1 2 3 4 5 6 7 8 9 10]

	// Print the capacity of the slice
	fmt.Println("Capacity:", cap(numbers)) // Outputs: Capacity: 12 (initial capacity is doubled when appended beyond the original capacity)

	// Access and print each element of the slice using a for loop
	fmt.Print("Elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Print(" ", numbers[i])
	}
	fmt.Println() // Print a newline after the elements
}
