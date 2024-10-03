package prac8

import (
	"errors"
	"fmt"
)

// Book struct represents a book with a title and author
type Book struct {
	Title  string
	Author string
}

// Library struct represents a collection of books
type Library struct {
	Books []Book
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

// FindBookByTitle searches for a book by its title
func (l *Library) FindBookByTitle(title string) (Book, error) {
	for _, book := range l.Books {
		if book.Title == title {
			return book, nil
		}
	}
	return Book{}, errors.New("book not found")
}

// ListBooks lists all books in the library
func (l *Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	fmt.Println("Books in the library:")
	for _, book := range l.Books {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}
}

// Run function to execute the practice code
func Run() {
	library := Library{}

	// Adding books to the library
	library.AddBook(Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan"})
	library.AddBook(Book{Title: "Go in Action", Author: "William Kennedy"})

	// Listing all books
	library.ListBooks()

	// Finding a book
	book, err := library.FindBookByTitle("The Go Programming Language")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found book: %s by %s\n", book.Title, book.Author)
	}

	// Trying to find a non-existent book
	book, err = library.FindBookByTitle("Unknown Book")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
