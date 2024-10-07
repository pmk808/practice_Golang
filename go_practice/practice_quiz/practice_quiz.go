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
		// Principles of Microservices
		{
			"What is the principle of Single Responsibility in microservices?",
			[]string{"a) Each service should do one thing well", "b) Services should be tightly coupled", "c) All services should share a database", "d) Services should be monolithic"},
			"a",
		},
		{
			"What does Autonomy mean in the context of microservices?",
			[]string{"a) Services are dependent on each other", "b) Services can be developed and deployed independently", "c) All services must use the same technology", "d) Services cannot be scaled"},
			"b",
		},
		{
			"What is meant by Decentralization in microservices?",
			[]string{"a) Centralized decision-making", "b) Avoiding single points of failure", "c) All data is stored in one place", "d) Services must be identical"},
			"b",
		},
		{
			"What does Isolation of Failures refer to in microservices?",
			[]string{"a) All services fail together", "b) Issues in one service shouldn't affect others", "c) Services must share resources", "d) Services are always online"},
			"b",
		},
		{
			"What is Continuous Delivery in microservices?",
			[]string{"a) Deploying all services at once", "b) Frequent, automated deployments of individual services", "c) Manual deployment of services", "d) Only deploying when necessary"},
			"b",
		},
		{
			"What does Evolutionary Design mean in microservices?",
			[]string{"a) Services cannot change", "b) Services can evolve independently", "c) All services must be rewritten", "d) Services are static"},
			"b",
		},

		// Advantages of Microservices
		{
			"What is a key advantage of microservices regarding scalability?",
			[]string{"a) Scale the entire application", "b) Scale individual components based on demand", "c) No scaling is possible", "d) All services must scale together"},
			"b",
		},
		{
			"What does Technology Diversity allow in microservices?",
			[]string{"a) Use of a single technology stack", "b) Freedom to use different technologies for different services", "c) All services must use the same programming language", "d) No technology choices"},
			"b",
		},
		{
			"How does microservices architecture improve development speed?",
			[]string{"a) Smaller codebases are easier to understand", "b) All code is in one place", "c) Larger teams are required", "d) Slower development due to complexity"},
			"a",
		},
		{
			"What is improved fault isolation in microservices?",
			[]string{"a) Failures affect all services", "b) Failures are contained within individual services", "c) All services must be online", "d) Services cannot fail"},
			"b",
		},
		{
			"What does it mean to innovate easily in microservices?",
			[]string{"a) All features must be tested together", "b) Experimenting with new features in isolated services", "c) No innovation is possible", "d) All services must be updated simultaneously"},
			"b",
		},

		// Challenges of Microservices
		{
			"What is a challenge of distributed system complexity in microservices?",
			[]string{"a) Easy inter-service communication", "b) Managing inter-service communication and data consistency", "c) All services are in one place", "d) No challenges exist"},
			"b",
		},
		{
			"What does operational overhead refer to in microservices?",
			[]string{"a) Fewer services to manage", "b) More services mean more components to monitor", "c) All services are automated", "d) No monitoring is needed"},
			"b",
		},
		{
			"What is a testing complexity challenge in microservices?",
			[]string{"a) Testing is simpler", "b) Need to test both individual services and their interactions", "c) No testing is required", "d) All tests are manual"},
			"b",
		},
		{
			"What does deployment complexity mean in microservices?",
			[]string{"a) Easy to deploy all services", "b) Coordinating updates across multiple services", "c) No deployment is needed", "d) All services are deployed at once"},
			"b",
		},
		{
			"What is network latency in microservices?",
			[]string{"a) Fast communication between services", "b) Communication between services adds overhead", "c) No latency exists", "d) All services are local"},
			"b",
		},

		// Comparison with Monolithic Architecture
		{
			"What is a key difference between monolithic and microservices architecture?",
			[]string{"a) Monoliths are easier to develop initially", "b) Microservices are easier to maintain as they grow", "c) Both are the same", "d) Monoliths cannot scale"},
			"a",
		},
		{
			"In terms of deployment, how do monolithic and microservices architectures differ?",
			[]string{"a) Monoliths deploy individual services", "b) Microservices deploy the entire application at once", "c) Monoliths deploy the entire application at once", "d) Microservices cannot be deployed"},
			"c",
		},
		{
			"What is a scaling difference between monolithic and microservices architectures?",
			[]string{"a) Monoliths scale individual components", "b) Microservices scale the entire application", "c) Monoliths scale the entire application", "d) Both scale the same way"},
			"c",
		},
		{
			"What is a technology stack difference between monolithic and microservices architectures?",
			[]string{"a) Monoliths can use multiple technologies", "b) Microservices can use different technologies for different services", "c) Both use the same technology stack", "d) Monoliths cannot use different technologies"},
			"b",
		},
		{
			"What is a team structure difference between monolithic and microservices architectures?",
			[]string{"a) Monoliths require smaller teams", "b) Microservices allow for smaller, specialized teams", "c) Both require large teams", "d) Monoliths cannot have specialized teams"},
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
