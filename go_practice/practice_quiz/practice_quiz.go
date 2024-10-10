package practice_quiz

import (
	"fmt"
)

func Run() {
	// Questions and answers for Service Discovery, Configuration Management, Logging, Monitoring, and Testing
	questions := []struct {
		question string
		options  []string
		answer   string
	}{
		// Service Discovery - Theory
		{
			"Which of the following is a key feature of service discovery?",
			[]string{"a) Caching", "b) Registration", "c) Version control", "d) Encryption"},
			"b",
		},
		{
			"In service discovery, what is the purpose of health checking?",
			[]string{"a) Ensure a service is functioning correctly", "b) Load balancing", "c) Caching data", "d) Configuring environment variables"},
			"a",
		},
		{
			"What does Consul primarily provide in a microservices architecture?",
			[]string{"a) Service discovery and health checking", "b) Authentication and authorization", "c) Database management", "d) Message queuing"},
			"a",
		},
		{
			"Which protocol is typically used by Consul for service discovery communication?",
			[]string{"a) HTTP", "b) TCP", "c) gRPC", "d) DNS"},
			"d",
		},
		{
			"What kind of data store is etcd?",
			[]string{"a) Distributed key-value store", "b) Relational database", "c) NoSQL document store", "d) In-memory cache"},
			"a",
		},
		{
			"In the food court analogy for service discovery, what does the 'central directory' represent?",
			[]string{"a) The menu", "b) The service registry", "c) The kitchen", "d) The customers"},
			"b",
		},

		// Configuration Management - Theory
		{
			"Which Go library is commonly used for managing application configuration?",
			[]string{"a) sqlx", "b) Viper", "c) Gin", "d) Echo"},
			"b",
		},
		{
			"In configuration management, what is the advantage of using environment variables?",
			[]string{"a) Centralized configuration", "b) Encryption", "c) Avoiding hardcoding sensitive data", "d) Improving performance"},
			"c",
		},
		{
			"Which of the following formats can Viper read from?",
			[]string{"a) JSON", "b) YAML", "c) Environment variables", "d) All of the above"},
			"d",
		},
		{
			"What is a best practice for handling configuration changes without restarting services?",
			[]string{"a) Hardcoding values", "b) Using environment variables", "c) Utilizing Viper's dynamic configuration reloading", "d) Embedding configurations in binary"},
			"c",
		},
		{
			"In the recipe book analogy for configuration management, what do 'recipes' represent?",
			[]string{"a) Microservices", "b) Configurations", "c) Databases", "d) APIs"},
			"b",
		},

		// Logging - Theory
		{
			"Why is structured logging preferred in microservices?",
			[]string{"a) Easier to read", "b) Allows more precise searching and filtering", "c) Takes less space", "d) Faster execution"},
			"b",
		},
		{
			"What Go package is commonly used for structured logging?",
			[]string{"a) log/slog", "b) fmt", "c) log", "d) Gin"},
			"a",
		},

		// Monitoring - Theory
		{
			"What is Prometheus primarily used for in microservices?",
			[]string{"a) Logging", "b) Metrics collection", "c) Distributed tracing", "d) Message queuing"},
			"b",
		},
		{
			"In Prometheus monitoring, what is a 'scrape'? ",
			[]string{"a) A query to the database", "b) A method to trace errors", "c) Collecting metrics from a target", "d) A method to log HTTP requests"},
			"c",
		},

		// Distributed Tracing - Theory
		{
			"Which tool is used for distributed tracing in microservices?",
			[]string{"a) Prometheus", "b) Grafana", "c) Jaeger", "d) Kubernetes"},
			"c",
		},
		{
			"Distributed tracing helps in tracking which of the following?",
			[]string{"a) User activity logs", "b) Request flows across services", "c) Metrics data", "d) Configuration changes"},
			"b",
		},

		// Testing Strategies - Theory
		{
			"What is the purpose of unit testing in microservices?",
			[]string{"a) To test individual services in isolation", "b) To test communication between services", "c) To test the entire system", "d) To test configuration changes"},
			"a",
		},
		{
			"Which Go package is commonly used for integration testing?",
			[]string{"a) log/slog", "b) net/http/httptest", "c) os", "d) fmt"},
			"b",
		},
		{
			"What type of testing focuses on the entire microservices ecosystem?",
			[]string{"a) Unit testing", "b) Integration testing", "c) End-to-end testing", "d) Stress testing"},
			"c",
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
