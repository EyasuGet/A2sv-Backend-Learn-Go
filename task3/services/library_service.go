package services

import (
	"errors"
	"library_management/models"
	"strconv"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct{
	Books map[int]models.Book
	Members map[int]*models.Member
}


func NewLibrary() *Library{
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]*models.Member),
	}
}

func (l *Library) AddBook(book models.Book){

	if book.Status == ""{
		book.Status = "Available"
	}

	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (lib *Library) BorrowBook(bookID int, memberID int) error{
	book, exists := lib.Books[bookID]

	if !exists{
		return errors.New("book not fouond")
	}
	if book.Status != "Available"{
		return errors.New("book is not available")
	}

	member, exists := lib.Members[memberID]
	if !exists {
		member = &models.Member{ID: memberID, Name: "member" + strconv.Itoa(memberID)}
		lib.Members[memberID] = member
	}

	book.Status = "Borrowed"
	lib.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	member, exists := lib.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			book := lib.Books[bookID]
			book.Status = "Available"
			lib.Books[bookID] = book
			return nil
		}
	}
	return errors.New("book not borrowed by this member")
}

func (l *Library) ListAvailableBooks() []models.Book{
	available := []models.Book{}
	for _, book := range l.Books{
		if book.Status == "Available"{
			available = append(available, book)
		}
	}

	return available
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := lib.Members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}