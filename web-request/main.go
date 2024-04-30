package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Learn Web Services")

	// Send a GET request to retrieve data from the provided URL
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	// Check if there was an error while sending the request
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	// Ensure that the response body is closed after the function returns
	defer res.Body.Close()

	// Read the response body data
	data, err := ioutil.ReadAll(res.Body)

	// Check if there was an error while reading the response body
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	// Print the response data
	fmt.Println("Response is:", string(data))
}
