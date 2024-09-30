package prac5

import (
	"fmt"
)

// Add function to add two integers
func Add(a int, b int) int {
	return a - b
}

// Run function to execute the practice code
func Run() {
	result := Add(3, 4)
	fmt.Println("Result of Add(3, 4):", result)
}
