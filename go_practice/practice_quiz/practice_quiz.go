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
		// Gin for HTTP Routing
		{
			"In the traffic controller analogy for Gin, what do the 'roads' represent?",
			[]string{"a) Databases", "b) Servers", "c) Endpoints", "d) Middleware"},
			"c",
		},
		{
			"Which of the following is a key feature of Gin?",
			[]string{"a) Built-in database", "b) Automatic code generation", "c) Middleware support", "d) Machine learning capabilities"},
			"c",
		},
		{
			"In a Gin router, what does `r.GET('/hello', handleHello)` do?",
			[]string{"a) Sends a GET request", "b) Creates a new endpoint", "c) Sets up a route to handle GET requests for /hello", "d) Defines a function named handleHello"},
			"c",
		},

		// gRPC for Inter-service Communication
		{
			"In the postal service analogy for gRPC, what do 'postcards with predefined fields' represent?",
			[]string{"a) JSON", "b) XML", "c) Protocol Buffers", "d) HTML"},
			"c",
		},
		{
			"What is a key advantage of using gRPC in microservices?",
			[]string{"a) It's easier to learn than REST", "b) It uses human-readable formats", "c) It's more efficient for data serialization", "d) It doesn't require any configuration"},
			"c",
		},
		{
			"In a .proto file for gRPC, what does `rpc SayHello (HelloRequest) returns (HelloReply) {}` define?",
			[]string{"a) A new data type", "b) A service method", "c) A variable", "d) A package"},
			"b",
		},

		// Docker for Containerization
		{
			"In the shipping analogy for Docker, what does a container represent?",
			[]string{"a) A server", "b) An application with its environment", "c) A programming language", "d) A database"},
			"b",
		},
		{
			"What is a key benefit of using Docker in microservices architecture?",
			[]string{"a) It makes the application run faster", "b) It provides a consistent environment across development and production", "c) It automatically scales the application", "d) It writes code for you"},
			"b",
		},
		{
			"In a Dockerfile, what does the `EXPOSE` instruction do?",
			[]string{"a) Opens a port on the host machine", "b) Installs network tools", "c) Informs Docker that the container listens on the specified port", "d) Creates a new network interface"},
			"c",
		},

		// Kubernetes for Orchestration
		{
			"In the logistics company analogy for Kubernetes, what do the 'shipping boxes' represent?",
			[]string{"a) Pods", "b) Nodes", "c) Services", "d) Containers"},
			"d",
		},
		{
			"What is a key feature of Kubernetes in managing microservices?",
			[]string{"a) It writes your application code", "b) It provides automatic scaling and self-healing", "c) It replaces the need for databases", "d) It makes your application faster"},
			"b",
		},
		{
			"In a Kubernetes deployment YAML, what does `replicas: 3` specify?",
			[]string{"a) The number of nodes in the cluster", "b) The number of containers to run", "c) The number of databases to create", "d) The number of networks to set up"},
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
