package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Introduction
	fmt.Println("Learning URL")

	// Example URL
	myURL := "http://www.example.com:8080/profile"

	// Printing the type of the URL string
	fmt.Printf("The Type of URL is: %T \n", myURL)

	// Parsing the URL string
	parsedURL, err := url.Parse(myURL)

	// Check for parsing errors
	if err != nil {
		// Print error message if parsing fails
		fmt.Println("Can't Parse URL")
	} else {
		// Print the type of the parsed URL
		fmt.Printf("The Type of URL is: %T \n", parsedURL)

		// Accessing different components of the parsed URL
		// Scheme
		fmt.Println("Scheme:", parsedURL.Scheme)

		// Host
		fmt.Println("Host:", parsedURL.Host)

		// Path
		fmt.Println("Path:", parsedURL.Path)
	}
}
