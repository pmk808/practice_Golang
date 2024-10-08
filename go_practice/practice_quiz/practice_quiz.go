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
		// RESTful API Design and Implementation
		{
			"In the library analogy for RESTful APIs, what do the 'shelves' represent?",
			[]string{"a) Resources", "b) HTTP methods", "c) Endpoints", "d) Data types"},
			"c",
		},
		{
			"Which HTTP method is analogous to 'updating information in a book' in a RESTful API?",
			[]string{"a) GET", "b) POST", "c) PUT/PATCH", "d) DELETE"},
			"c",
		},
		{
			"In a Go RESTful API using Gin, what does `router.GET('/books', getBooks)` do?",
			[]string{"a) Defines a function named getBooks", "b) Creates a new book", "c) Sets up a route to handle GET requests for /books", "d) Deletes all books"},
			"c",
		},
		{
			"What's the purpose of `c.JSON(200, books)` in a Gin handler function?",
			[]string{"a) To create a new JSON file", "b) To send a JSON response with status code 200", "c) To parse incoming JSON", "d) To validate JSON data"},
			"b",
		},

		// Goroutines and Concurrency
		{
			"In the kitchen analogy for goroutines, what represents a goroutine?",
			[]string{"a) A dish", "b) The kitchen", "c) A chef", "d) A recipe"},
			"c",
		},
		{
			"What does the `go` keyword do in Go?",
			[]string{"a) Creates a new variable", "b) Defines a new function", "c) Starts a new goroutine", "d) Imports a package"},
			"c",
		},
		{
			"Why might we use `time.Sleep()` in a program with goroutines?",
			[]string{"a) To make the program faster", "b) To allow time for goroutines to complete", "c) To create more goroutines", "d) To stop all goroutines"},
			"b",
		},
		{
			"What's a key advantage of using goroutines in a microservice?",
			[]string{"a) They make the code easier to read", "b) They allow efficient handling of multiple tasks", "c) They automatically fix bugs", "d) They reduce the need for testing"},
			"b",
		},

		// Channels for Inter-service Communication
		{
			"In the factory analogy, what do channels represent?",
			[]string{"a) Workers", "b) Products", "c) Conveyor belts", "d) Managers"},
			"c",
		},
		{
			"What's the primary purpose of channels in Go?",
			[]string{"a) To store data permanently", "b) To create new goroutines", "c) To allow safe communication between goroutines", "d) To replace functions"},
			"c",
		},
		{
			"How do you create a new channel in Go?",
			[]string{"a) new(chan)", "b) make(chan Type)", "c) channel.New()", "d) createChannel()"},
			"b",
		},
		{
			"What does `for order := range orders` do in a Go program?",
			[]string{"a) Creates new orders", "b) Deletes all orders", "c) Receives orders from a channel until it's closed", "d) Sends orders to a channel"},
			"c",
		},

		// Context Package for Request Scoping and Cancellation
		{
			"In the mission control analogy, what does the context represent?",
			[]string{"a) The entire mission", "b) The control panel with abort button", "c) The mission commander", "d) The spaceship"},
			"b",
		},
		{
			"What's a primary use of the context package in Go microservices?",
			[]string{"a) To create new services", "b) To store permanent data", "c) To manage request lifecycles and cancellation", "d) To replace databases"},
			"c",
		},
		{
			"What does `ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)` do?",
			[]string{"a) Cancels the context immediately", "b) Creates a context that will timeout after 5 seconds", "c) Pauses the program for 5 seconds", "d) Creates 5 new goroutines"},
			"b",
		},
		{
			"What happens when a context is canceled in a Go program?",
			[]string{"a) The entire program crashes", "b) All goroutines stop immediately", "c) The database is wiped", "d) Functions can check for cancellation and stop gracefully"},
			"d",
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
