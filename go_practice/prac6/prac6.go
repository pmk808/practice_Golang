package prac6

import (
	"fmt"
)

// Run function to execute the practice code
func Run() {
	// Example of multiple variables that are declared but not used
	var unusedVariable int
	var anotherUnusedVariable string // Deliberately unused variable

	fmt.Println(unusedVariable)
	fmt.Println(anotherUnusedVariable)
	// Example of a potential issue: using a loop variable correctly
	for i := 0; i < 5; i++ {
		fmt.Println("Current value:", i)
	}

	// Example of a function that returns an error but is not checked
	result, err := divide(10, 0) // This will cause an error
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Deliberate formatting issue
	fmt.Println("This line has inconsistent spacing.") //  <-- Extra space here
}

// divide function that returns an error if division by zero is attempted
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero") // This is correct, but we won't check it properly
	}
	return a / b, nil
}
