package prac2

import (
	"fmt"
	"math/rand"
	"time"
)

// Custom error type
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// Function that may return an error
func generateNumber(max int) (int, error) {
	if max <= 0 {
		return 0, &ValidationError{Message: "Max must be greater than 0"}
	}
	return rand.Intn(max), nil
}

// Function to be run as a goroutine
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		results <- j * 2
	}
}

func Run() {
	// Set up channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs and handle potential errors
	for i := 1; i <= 5; i++ {
		num, err := generateNumber(10)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		jobs <- num
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}

	// Demonstrate defer
	defer fmt.Println("Main function is about to exit")

	fmt.Println("All jobs completed")
}
