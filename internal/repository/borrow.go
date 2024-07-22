package repository

import (
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/utils"
)

type BorrowRepository struct {
	DB map[string]*domain.Borrow
}

func NewBorrowRepository() IBorrowRepository {
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

func (r *BorrowRepository) Borrow(borrow *domain.Borrow) (*domain.Borrow, error) {
	for _, borrowedBook := range r.DB {
		if borrowedBook.BookID == borrow.BookID &&
			borrowedBook.Status == domain.StatusBorrowed &&
			utils.TimeIsBetween(borrow.StartDate, borrowedBook.StartDate, borrowedBook.EndDate) {
			return nil, utils.ErrBookCurrentlyBorrowed
		}
	}

	r.DB[borrow.ID] = borrow
	return borrow, nil
}

func (r *BorrowRepository) GetBorrowedBook(userID, bookID string) (*domain.Borrow, error) {
	for _, borrowedBook := range r.DB {
		if borrowedBook.UserID == userID &&
			borrowedBook.BookID == bookID &&
			borrowedBook.Status == domain.StatusBorrowed {
			return borrowedBook, nil
		}
	}

	return nil, utils.ErrBorrowedBookNotFound
}

func (r *BorrowRepository) UpdateStatus(borrowID string, status domain.Status) (*domain.Borrow, error) {
	if borrow, exists := r.DB[borrowID]; exists {
		borrow.Status = status
		r.DB[borrowID] = borrow
		return borrow, nil
	}

	return nil, utils.ErrBorrowedBookNotFound
}
