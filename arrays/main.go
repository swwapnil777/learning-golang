package main

import "fmt"

func main() {
	// Displaying a message to indicate the topic
	fmt.Println("Learning Arrays")

	// Declaring and initializing an array of strings
	var names [5]string
	names[0] = "Swapnil"
	names[1] = "Ritesh"
	names[2] = "Pratik"
	names[3] = "Atul"
	names[4] = "Kaustub"

	// Printing the entire names array
	fmt.Println("The names are:", names)

	// Printing the length of the names array
	fmt.Println("The length of array is:", len(names))

	// Printing the name at 0th index
	fmt.Println("The 0th Index name is:", names[0])

	// Demonstrating additional ways to print array elements
	fmt.Println("Individual names:")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Creating and initializing an array of integers
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	// Printing the entire numbers array
	fmt.Println("The numbers are:", numbers)

	// Creating and initializing an array of booleans
	var flags [5]bool
	flags[0] = true
	flags[1] = false
	flags[2] = true
	flags[3] = false
	flags[4] = true

	// Printing the entire flags array
	fmt.Println("The flags are:", flags)

	// Creating and initializing an array of floats
	var grades [5]float64
	grades[0] = 92.5
	grades[1] = 85.0
	grades[2] = 78.5
	grades[3] = 90.0
	grades[4] = 87.5

	// Printing the entire grades array
	fmt.Println("The grades are:", grades)
}
