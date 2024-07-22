package repository

import (
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/utils"
)

type BookRepository struct {
	db map[string]*domain.Book
}

func NewBookRepository() IBookRepository {
	return &BookRepository{
		db: make(map[string]*domain.Book),
	}
}

func (r *BookRepository) GetAll() []*domain.Book {
	books := make([]*domain.Book, 0)
	for _, book := range r.db {
		books = append(books, book)
	}
	return books
}

func (r *BookRepository) Add(book *domain.Book) *domain.Book {
	r.db[book.ID] = book
	return book
}

func (r *BookRepository) GetByID(bookID string) (*domain.Book, error) {
	if book, exists := r.db[bookID]; exists {
		return book, nil
	}

	return nil, utils.ErrBookNotFound
}
