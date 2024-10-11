package practice_quiz

import (
	"fmt"
)

func Run() {
	questions := []struct {
		question string
		options  []string
		answer   string
	}{
		// Security Considerations - Authentication and Authorization
		{
			"What does JWT stand for in the context of authentication?",
			[]string{"a) Java Web Token", "b) JSON Web Token", "c) JavaScript Web Token", "d) Joint Web Technology"},
			"b",
		},
		{
			"In the backstage pass analogy for JWT, what does the token represent?",
			[]string{"a) The security guard", "b) The concert venue", "c) The backstage pass", "d) The ticket office"},
			"c",
		},
		{
			"Which of the following is a key advantage of using JWT for authentication?",
			[]string{"a) Encryption of all data", "b) Stateless authentication", "c) Automatic password reset", "d) Two-factor authentication"},
			"b",
		},

		// Security Considerations - HTTPS
		{
			"In the analogy for HTTPS, what does the locked briefcase represent?",
			[]string{"a) The server", "b) The encrypted message", "c) The recipient", "d) The internet"},
			"b",
		},
		{
			"What is the main purpose of using HTTPS in microservices communication?",
			[]string{"a) Faster data transfer", "b) Reduced server load", "c) Encrypted data transmission", "d) Improved SEO"},
			"c",
		},

		// Security Considerations - Input Validation and Sanitization
		{
			"In the club bouncer analogy for input validation, what does the bouncer represent?",
			[]string{"a) The database", "b) The validator", "c) The user", "d) The server"},
			"b",
		},
		{
			"Which of the following is NOT a typical goal of input sanitization?",
			[]string{"a) Prevent XSS attacks", "b) Ensure data consistency", "c) Encrypt all input data", "d) Remove potentially harmful characters"},
			"c",
		},

		// Deployment and CI/CD - CI/CD Pipeline
		{
			"In the car factory analogy for CI/CD, what do the quality check stations represent?",
			[]string{"a) Code reviews", "b) Pipeline stages", "c) Deployment servers", "d) End users"},
			"b",
		},
		{
			"What is the primary purpose of a CI/CD pipeline in microservices development?",
			[]string{"a) To increase development time", "b) To automate testing and deployment processes", "c) To replace manual coding", "d) To reduce the number of features"},
			"b",
		},

		// Deployment and CI/CD - Blue-Green Deployment
		{
			"In the highway analogy for Blue-Green deployment, what does switching traffic represent?",
			[]string{"a) Load balancing", "b) Redirecting user requests to the new version", "c) Shutting down old servers", "d) Testing in production"},
			"b",
		},
		{
			"What is a key advantage of Blue-Green deployment?",
			[]string{"a) Reduced development time", "b) Improved security", "c) Zero-downtime updates", "d) Automatic bug fixing"},
			"c",
		},

		// Deployment and CI/CD - Canary Deployment
		{
			"In the coal mine analogy for Canary deployment, what does the canary represent?",
			[]string{"a) The production environment", "b) The new version of the software", "c) The end users", "d) The old version of the software"},
			"b",
		},
		{
			"What is the main benefit of using Canary deployment?",
			[]string{"a) Faster development", "b) Reduced server costs", "c) Gradual rollout with reduced risk", "d) Automatic error correction"},
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
