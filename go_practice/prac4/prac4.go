package prac4

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func demoFmtIO() {
	// Print to standard output
	fmt.Println("Demonstrating fmt package for I/O:")

	// Print formatted string
	name := "Gopher"
	age := 10
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", name, age)

	// Read from standard input
	fmt.Print("Enter your favorite programming language: ")
	var lang string
	fmt.Scanln(&lang)
	fmt.Printf("You entered: %s\n", lang)

	// Write to a file using fmt
	file, err := os.Create("fmt_output.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "This is a line written to a file using fmt.Fprintln")
	fmt.Fprintf(file, "We can also use formatting: %s is %d years old\n", name, age)

	fmt.Println("Check fmt_output.txt for the file output.")
}

// readFileContent reads the content of a file and returns it as a string
func readFileContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// handleFileContent serves the content of a file over HTTP
func handleFileContent(w http.ResponseWriter, r *http.Request) {
	filename := "./files/example.txt" // You'll need to create this file
	content, err := readFileContent(filename)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File content:\n%s", content)
}

// loggingMiddleware logs information about incoming requests
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func Run() {
	demoFmtIO()
	// Create a file for demonstration
	filename := "example.txt"
	content := "Hello, this is a sample file for our Go standard library practice!"
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	// Set up HTTP server
	http.HandleFunc("/", loggingMiddleware(handleFileContent))

	fmt.Println("Server is running on http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop the server")

	// Start the server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
