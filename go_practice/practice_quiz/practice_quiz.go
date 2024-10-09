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
		// Database Integration
		{
			"In the library analogy for database integration, what does the 'librarian' represent?",
			[]string{"a) The database", "b) The user", "c) The microservice", "d) The data"},
			"c",
		},
		{
			"Which package in Go provides a consistent interface for SQL databases?",
			[]string{"a) sql.DB", "b) database/sql", "c) sqlx", "d) gorm"},
			"b",
		},
		{
			"In MongoDB, what is the equivalent of an SQL table?",
			[]string{"a) Document", "b) Collection", "c) Database", "d) Field"},
			"b",
		},

		// SQL vs NoSQL
		{
			"In the library analogy, what do 'books stored on shelves with a specific structure' represent?",
			[]string{"a) NoSQL databases", "b) SQL databases", "c) In-memory caches", "d) File systems"},
			"b",
		},
		{
			"Which of the following is a characteristic of NoSQL databases?",
			[]string{"a) Strict schema", "b) ACID transactions", "c) Flexible schema", "d) SQL query language"},
			"c",
		},
		{
			"What does BSON stand for in the context of MongoDB?",
			[]string{"a) Basic SQL Object Notation", "b) Binary JSON", "c) Boolean Serialized Object Notation", "d) Baselined Serialized Object Notation"},
			"b",
		},

		// Service Discovery
		{
			"In the food court analogy for service discovery, what does the 'central directory' represent?",
			[]string{"a) The menu", "b) The service registry", "c) The kitchen", "d) The cash register"},
			"b",
		},
		{
			"Which of the following is NOT a key concept in service discovery?",
			[]string{"a) Registration", "b) Discovery", "c) Health Checking", "d) Load Balancing"},
			"d",
		},
		{
			"What is the primary purpose of Consul in a microservices architecture?",
			[]string{"a) Database management", "b) Service discovery and health checking", "c) Code deployment", "d) User authentication"},
			"b",
		},

		// Configuration Management
		{
			"In the recipe book analogy for configuration management, what do the 'recipes' represent?",
			[]string{"a) Services", "b) Databases", "c) Configurations", "d) APIs"},
			"c",
		},
		{
			"Which of the following is a feature of the Viper configuration management library?",
			[]string{"a) Automatic code generation", "b) Database migrations", "c) Reading from multiple configuration formats", "d) Load balancing"},
			"c",
		},
		{
			"What is a best practice for handling sensitive information in configuration management?",
			[]string{"a) Hardcoding in the application", "b) Storing in version control", "c) Using environment variables", "d) Sharing through public APIs"},
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
