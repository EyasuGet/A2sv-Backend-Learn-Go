package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func StartConsole(library services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nWelcome to My Library")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		option := strings.TrimSpace(input)

		switch option {
		case "1":
			addBook(reader, library)
		case "2":
			removeBook(reader, library)
		case "3":
			borrowBook(reader, library)
		case "4":
			returnBook(reader, library)
		case "5":
			listAvailableBooks(library)
		case "6":
			listBorrowedBooks(reader, library)
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func addBook(reader *bufio.Reader, library services.LibraryManager) {
	fmt.Print("Enter Book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter Book Title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter Book Author: ")
	author, _ := reader.ReadString('\n')

	book := models.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
	}

	library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func removeBook(reader *bufio.Reader, library services.LibraryManager) {
	fmt.Print("Enter Book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	library.RemoveBook(id)
	fmt.Println("Book removed.")
}

func borrowBook(reader *bufio.Reader, library services.LibraryManager) {
	fmt.Print("Enter Book ID to borrow: ")
	bookStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookStr))

	fmt.Print("Enter Member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberStr))

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func returnBook(reader *bufio.Reader, library services.LibraryManager) {
	fmt.Print("Enter Book ID to return: ")
	bookStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookStr))

	fmt.Print("Enter Member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberStr))

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func listAvailableBooks(library services.LibraryManager) {
	books := library.ListAvailableBooks()
	fmt.Println("\nAvailable Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func listBorrowedBooks(reader *bufio.Reader, library services.LibraryManager) {
	fmt.Print("Enter Member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberStr))

	books := library.ListBorrowedBooks(memberID)
	fmt.Println("\nBorrowed Books:")
	if len(books) == 0 {
		fmt.Println("This member has not borrowed any books.")
		return
	}
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}