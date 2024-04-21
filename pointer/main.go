package main

import "fmt"

// modifyByReference takes a pointer to an integer and adds 10 to the value it points to.
func modifyByReference(value *int) {
	*value += 10
}

func main() {
	// Print a greeting message.
	fmt.Println("Hello Golang")

	// Declare and initialize an integer variable num with a value of 10.
	var num int = 10

	// Declare a pointer variable ptr and assign it the address of num.
	var ptr *int = &num

	// Print the address stored in ptr.
	fmt.Println("The Address is:", ptr)

	// Print the value that ptr points to.
	fmt.Println("The Value of num is:", *ptr)

	// Declare a nil pointer variable pointer.
	var pointer *int

	// Check if pointer is nil and print a message accordingly.
	if pointer == nil {
		fmt.Println("Pointer is Not Assigned")
	}

	// Call modifyByReference function with the address of num as argument to modify its value.
	modifyByReference(&num)

	// Print the modified value of num.
	fmt.Println("Value is Changed", num)
}
