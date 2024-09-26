package practice_quiz

import (
	"fmt"
)

func Run() {
	// Questions and answers
	questions := []struct {
		question string
		options  []string
		answer   string
	}{
		{
			"What is the purpose of the 'fmt' package in Go?",
			[]string{"a) Formatting I/O", "b) File handling", "c) Networking", "d) Concurrency"},
			"a",
		},
		{
			"What does 'defer' do in Go?",
			[]string{"a) Delays function execution", "b) Stops the program", "c) Creates a new goroutine", "d) Returns a value"},
			"a",
		},
		{
			"What is a goroutine?",
			[]string{"a) A type of variable", "b) A lightweight thread", "c) A function", "d) A package"},
			"b",
		},
		{
			"What is the purpose of channels in Go?",
			[]string{"a) To store data", "b) To communicate between goroutines", "c) To handle errors", "d) To format strings"},
			"b",
		},
		{
			"What does the 'os' package allow you to do?",
			[]string{"a) Handle HTTP requests", "b) Interact with the operating system", "c) Format output", "d) Manage concurrency"},
			"b",
		},
		{
			"What is a struct in Go?",
			[]string{"a) A built-in data type", "b) A collection of fields", "c) A function type", "d) A package"},
			"b",
		},
		{
			"What does the 'math' package provide?",
			[]string{"a) String manipulation", "b) Mathematical functions", "c) File handling", "d) Network communication"},
			"b",
		},
		{
			"What is the output of 'fmt.Println(2 + 2)'?",
			[]string{"a) 2", "b) 4", "c) 22", "d) Error"},
			"b",
		},
		{
			"What is the purpose of the 'http' package?",
			[]string{"a) To handle file I/O", "b) To create web servers and clients", "c) To manage goroutines", "d) To format strings"},
			"b",
		},
		{
			"What does 'http.HandleFunc' do?",
			[]string{"a) Starts a web server", "b) Handles HTTP requests", "c) Defines a route", "d) All of the above"},
			"d",
		},
		{
			"What is the return type of 'os.ReadFile'?",
			[]string{"a) string", "b) []byte", "c) error", "d) int"},
			"b",
		},
		{
			"What is the purpose of the 'log' package?",
			[]string{"a) To format strings", "b) To log messages and errors", "c) To handle HTTP requests", "d) To manage concurrency"},
			"b",
		},
		{
			"What does 'fmt.Fprintf' do?",
			[]string{"a) Formats and prints to a file", "b) Reads from a file", "c) Writes to the console", "d) None of the above"},
			"a",
		},
		{
			"What is a common use case for 'defer'?",
			[]string{"a) To delay execution", "b) To ensure resources are released", "c) To create goroutines", "d) To format output"},
			"b",
		},
		{
			"What is the output of 'fmt.Sprintf'?",
			[]string{"a) Prints to the console", "b) Returns a formatted string", "c) Writes to a file", "d) None of the above"},
			"b",
		},
		{
			"What is the purpose of the 'ValidationError' struct in prac2.go?",
			[]string{"a) To handle HTTP errors", "b) To validate user input", "c) To represent custom errors", "d) To format strings"},
			"c",
		},
		{
			"What does 'os.Create' do?",
			[]string{"a) Creates a new file", "b) Opens an existing file", "c) Deletes a file", "d) Reads a file"},
			"a",
		},
		{
			"What is the purpose of the 'worker' function in prac2.go?",
			[]string{"a) To generate random numbers", "b) To process jobs concurrently", "c) To handle HTTP requests", "d) To read files"},
			"b",
		},
		{
			"What does 'close(jobs)' do in prac2.go?",
			[]string{"a) Closes the program", "b) Stops all goroutines", "c) Signals that no more jobs will be sent", "d) None of the above"},
			"c",
		},
		{
			"What is the output of 'fmt.Println'?",
			[]string{"a) Writes to a file", "b) Prints to the console", "c) Returns a string", "d) None of the above"},
			"b",
		},
	}

	score := 0

	// Loop through questions
	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.question)
		for _, option := range q.options {
			fmt.Println(option)
		}
		var answer string
		fmt.Print("Your answer (a/b/c/d): ")
		fmt.Scanln(&answer)

		if answer == q.answer {
			fmt.Println("Correct!\n")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer was %s.\n\n", q.answer)
		}
	}

	// Display final score
	fmt.Printf("Your final score is: %d out of %d\n", score, len(questions))
}