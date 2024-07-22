package repository

import "github.com/nadiannis/libry-api/internal/domain"

type BorrowRepository struct {
	DB map[string]*domain.Borrow
}

func NewBorrowRepository() *BorrowRepository {
	return &BorrowRepository{
		DB: make(map[string]*domain.Borrow),
	}
}

func (r *BorrowRepository) GetAll() []*domain.Borrow {
	borrowedBooks := make([]*domain.Borrow, 0)
	for _, borrowedBook := range r.DB {
		borrowedBooks = append(borrowedBooks, borrowedBook)
	}
	return borrowedBooks
}
