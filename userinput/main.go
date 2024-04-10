package main

import "fmt"

func main() {
	// Declare variables to store the user's name, age, and favorite color
	var name string
	var age int
	var color string

	// Ask the user for their name
	fmt.Println("Hey.. What's Your Name?")
	// Scan the user's input and store it in the 'name' variable
	fmt.Scan(&name)

	// Ask the user for their age
	fmt.Println("Nice to meet you,", name+". How old are you?")
	// Scan the user's input and store it in the 'age' variable
	fmt.Scan(&age)

	// Ask the user for their favorite color
	fmt.Println("Cool! What is your favorite color?")
	// Scan the user's input and store it in the 'color' variable
	fmt.Scan(&color)

	// Print the gathered information
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Favorite Color:", color)

	// Display a personalized message to the user
	fmt.Printf("Nice to meet you, %s! It's great that you're %d years old and your favorite color is %s.\n", name, age, color)
}
