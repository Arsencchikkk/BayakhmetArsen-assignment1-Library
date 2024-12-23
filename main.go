package main

import (
	"fmt"
	"github.com/Arsencchikkk/BayakhmetArsen-assignment1-Library/Library"
)

func main() {
	library := Library.NewLibrary()

	library.AddBook(Library.Book{ID: "001", Title: "Harry Potter", Author: "J.K. Rowling", IsBorrowed: false})
	library.AddBook(Library.Book{ID: "002", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", IsBorrowed: false})
	library.AddBook(Library.Book{ID: "003", Title: "1984", Author: "George Orwell", IsBorrowed: false})

	for {
		fmt.Println("\nLibrary menu:")
		fmt.Println("1. Add book")
		fmt.Println("2. Borrow book")
		fmt.Println("3. Return book")
		fmt.Println("4. See available books")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter book author: ")
			fmt.Scan(&author)
			library.AddBook(Library.Book{ID: id, Title: title, Author: author, IsBorrowed: false})
		case 2:
			var id string
			fmt.Print("Enter the book ID to borrow: ")
			fmt.Scan(&id)
			library.BorrowBook(id)
		case 3:
			var id string
			fmt.Print("Enter the book ID to return: ")
			fmt.Scan(&id)
			library.ReturnBook(id)
		case 4:
			fmt.Println("Available books:")
			library.ListBooks()
		case 5:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
