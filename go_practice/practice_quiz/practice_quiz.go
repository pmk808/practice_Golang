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
		// Basic Syntax and Concepts
		{
			"What is the zero value of an int in Go?",
			[]string{"a) 0", "b) -1", "c) 1", "d) Undefined"},
			"a",
		},
		{
			"Which control structure is used to execute code based on a condition?",
			[]string{"a) for", "b) if", "c) switch", "d) both b and c"},
			"d",
		},
		{
			"What keyword is used to define a function in Go?",
			[]string{"a) func", "b) define", "c) function", "d) method"},
			"a",
		},
		{
			"What is the purpose of the 'import' statement?",
			[]string{"a) To include packages", "b) To define variables", "c) To create functions", "d) To handle errors"},
			"a",
		},

		// Go-Specific Features
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
			"What does the 'defer' keyword do?",
			[]string{"a) Delays function execution", "b) Stops the program", "c) Creates a new goroutine", "d) Returns a value"},
			"a",
		},
		{
			"What is the purpose of the 'panic' function?",
			[]string{"a) To handle errors", "b) To stop the program", "c) To recover from a panic", "d) To log messages"},
			"b",
		},

		// Data Structures
		{
			"What is the zero value of a slice in Go?",
			[]string{"a) nil", "b) []", "c) 0", "d) Undefined"},
			"a",
		},
		{
			"What is a map in Go?",
			[]string{"a) A collection of key-value pairs", "b) A type of array", "c) A function", "d) A package"},
			"a",
		},
		{
			"What is the purpose of structs in Go?",
			[]string{"a) To define functions", "b) To group related data", "c) To handle errors", "d) To create goroutines"},
			"b",
		},
		{
			"What is the difference between an array and a slice?",
			[]string{"a) Arrays are fixed size, slices are dynamic", "b) Slices are fixed size, arrays are dynamic", "c) They are the same", "d) None of the above"},
			"a",
		},

		// Standard Library
		{
			"What does the 'fmt' package do?",
			[]string{"a) Formatting I/O", "b) File handling", "c) Networking", "d) Concurrency"},
			"a",
		},
		{
			"What is the purpose of the 'os' package?",
			[]string{"a) To handle HTTP requests", "b) To interact with the operating system", "c) To format output", "d) To manage concurrency"},
			"b",
		},
		{
			"What package is used for creating web servers?",
			[]string{"a) net/http", "b) fmt", "c) os", "d) time"},
			"a",
		},
		{
			"What function is used to read a file in Go?",
			[]string{"a) os.ReadFile", "b) fmt.ReadFile", "c) io.ReadFile", "d) file.Read"},
			"a",
		},

		// Testing
		{
			"What is the purpose of unit tests?",
			[]string{"a) To test individual components", "b) To test the entire application", "c) To format code", "d) To handle errors"},
			"a",
		},
		{
			"What is a table-driven test?",
			[]string{"a) A test with multiple cases", "b) A test that uses tables", "c) A test for web applications", "d) A test for databases"},
			"a",
		},

		// Tools
		{
			"What does 'go fmt' do?",
			[]string{"a) Formats Go code", "b) Compiles Go code", "c) Runs tests", "d) Creates modules"},
			"a",
		},
		{
			"What does 'go vet' do?",
			[]string{"a) Detects suspicious code", "b) Formats code", "c) Compiles code", "d) Runs the application"},
			"a",
		},

		// Project Structure
		{
			"What is the purpose of Go modules?",
			[]string{"a) To manage dependencies", "b) To format code", "c) To create tests", "d) To handle errors"},
			"a",
		},
		{
			"What is a common practice for organizing code in a Go project?",
			[]string{"a) Grouping by functionality", "b) Grouping by file size", "c) Random organization", "d) Grouping by author"},
			"a",
		},

		// Best Practices
		{
			"What is a guideline for writing effective Go code?",
			[]string{"a) Use clear naming conventions", "b) Avoid comments", "c) Write long functions", "d) Use global variables"},
			"a",
		},
		{
			"What is a common idiom in Go for handling errors?",
			[]string{"a) Ignore errors", "b) Return errors", "c) Log errors", "d) Panic on errors"},
			"b",
		},
	}

	// Quiz logic
	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.question)
		for _, option := range q.options {
			fmt.Println(option)
		}
		var answer string
		fmt.Print("Your answer: ")
		fmt.Scan(&answer)

		if answer == q.answer {
			fmt.Println("Correct!\n")
		} else {
			fmt.Printf("Wrong! The correct answer is: %s\n\n", q.answer)
		}
	}
}
