package repository

import "github.com/nadiannis/libry-api/internal/domain"

type BookRepository struct {
	DB map[string]*domain.Book
}

func (r *BookRepository) GetAll() []*domain.Book {
	books := make([]*domain.Book, 0)
	for _, book := range r.DB {
		books = append(books, book)
	}
	return books
}

func (r *BookRepository) Add(book *domain.Book) *domain.Book {
	r.DB[book.ID] = book
	return book
}
