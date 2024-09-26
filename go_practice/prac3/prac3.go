package prac3

import (
	"fmt"
)

// Book struct to represent a book
type Book struct {
	Title  string
	Author string
	Year   int
}

// Library struct to represent a library
type Library struct {
	Name  string
	Books []Book
}

// addBook adds a book to the library
func (l *Library) addBook(b Book) {
	l.Books = append(l.Books, b)
}

// findBooksByAuthor returns all books by a given author
func (l *Library) findBooksByAuthor(author string) []Book {
	var result []Book
	for _, book := range l.Books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

func Run() {
	// Array example
	var bestsellers [3]string
	bestsellers[0] = "The Go Programming Language"
	bestsellers[1] = "Go in Action"
	bestsellers[2] = "Concurrency in Go"

	fmt.Println("Bestsellers (Array):")
	for i, book := range bestsellers {
		fmt.Printf("%d. %s\n", i+1, book)
	}

	// Slice example
	genres := []string{"Fiction", "Non-fiction", "Science", "Programming"}
	fmt.Println("\nGenres (Slice):")
	for _, genre := range genres {
		fmt.Println("-", genre)
	}

	// Map example
	bookRatings := map[string]float64{
		"The Go Programming Language": 4.5,
		"Go in Action":                4.2,
		"Concurrency in Go":           4.7,
	}

	fmt.Println("\nBook Ratings (Map):")
	for book, rating := range bookRatings {
		fmt.Printf("%s: %.1f\n", book, rating)
	}

	// Struct and methods example
	myLibrary := Library{Name: "My Go Library"}

	myLibrary.addBook(Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Year: 2015})
	myLibrary.addBook(Book{Title: "Go in Action", Author: "William Kennedy", Year: 2015})
	myLibrary.addBook(Book{Title: "Concurrency in Go", Author: "Katherine Cox-Buday", Year: 2017})
	myLibrary.addBook(Book{Title: "Go Web Programming", Author: "Sau Sheong Chang", Year: 2016})

	fmt.Println("\nMy Library:")
	for _, book := range myLibrary.Books {
		fmt.Printf("- %s by %s (%d)\n", book.Title, book.Author, book.Year)
	}

	author := "William Kennedy"
	authorBooks := myLibrary.findBooksByAuthor(author)
	fmt.Printf("\nBooks by %s:\n", author)
	for _, book := range authorBooks {
		fmt.Printf("- %s (%d)\n", book.Title, book.Year)
	}
}
