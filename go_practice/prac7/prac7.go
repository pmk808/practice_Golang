package prac7

import (
	"fmt"
)

// Run function to execute the practice code
func Run() {
	fmt.Println("Welcome to Prac7!")
	
	// Example of using a simple function
	result := add(5, 10)
	fmt.Printf("The sum of 5 and 10 is: %d\n", result)
}

// add function to demonstrate a simple operation
func add(a int, b int) int {
	return a + b
}
