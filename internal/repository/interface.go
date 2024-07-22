package repository

import "github.com/nadiannis/libry-api/internal/domain"

type IUserRepository interface {
	GetAll() []*domain.User
	Add(user *domain.User) *domain.User
	GetByID(userID string) (*domain.User, error)
	AddBook(userID string, book *domain.Book) (*domain.Book, error)
	DeleteBookByID(userID, bookID string) error
}

type IBookRepository interface {
	GetAll() []*domain.Book
	Add(book *domain.Book) *domain.Book
	GetByID(bookID string) (*domain.Book, error)
}

type IBorrowRepository interface {
	GetAll() []*domain.Borrow
	Borrow(borrow *domain.Borrow) (*domain.Borrow, error)
	GetBorrowedBook(userID, bookID string) (*domain.Borrow, error)
	UpdateStatus(borrowID string, status domain.Status) (*domain.Borrow, error)
}
