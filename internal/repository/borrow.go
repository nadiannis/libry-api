package repository

import (
	"time"

	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/utils"
)

type BorrowRepository struct {
	db map[string]*domain.Borrow
}

func NewBorrowRepository() IBorrowRepository {
	return &BorrowRepository{
		db: make(map[string]*domain.Borrow),
	}
}

func (r *BorrowRepository) GetAll() []*domain.Borrow {
	borrowedBooks := make([]*domain.Borrow, 0)
	for _, borrowedBook := range r.db {
		borrowedBooks = append(borrowedBooks, borrowedBook)
	}
	return borrowedBooks
}

func (r *BorrowRepository) Borrow(borrow *domain.Borrow) (*domain.Borrow, error) {
	for _, borrowedBook := range r.db {
		if borrowedBook.BookID == borrow.BookID &&
			borrowedBook.Status == domain.StatusBorrowed &&
			utils.TimeIsBetween(borrow.StartDate, borrowedBook.StartDate, borrowedBook.EndDate) {
			return nil, utils.ErrBookCurrentlyBorrowed
		}
	}

	r.db[borrow.ID] = borrow
	return borrow, nil
}

func (r *BorrowRepository) GetBorrowedBook(userID, bookID string) (*domain.Borrow, error) {
	for _, borrowedBook := range r.db {
		if borrowedBook.UserID == userID &&
			borrowedBook.BookID == bookID &&
			borrowedBook.Status == domain.StatusBorrowed {
			return borrowedBook, nil
		}
	}

	return nil, utils.ErrBorrowedBookNotFound
}

func (r *BorrowRepository) UpdateStatus(borrowID string, status domain.Status) (*domain.Borrow, error) {
	if borrow, exists := r.db[borrowID]; exists {
		borrow.Status = status
		r.db[borrowID] = borrow
		return borrow, nil
	}

	return nil, utils.ErrBorrowedBookNotFound
}

func (r *BorrowRepository) UpdateDates(borrowID string, startDate, endDate time.Time) (*domain.Borrow, error) {
	if borrow, exists := r.db[borrowID]; exists {
		if startDate.After(endDate) {
			startDate, endDate = endDate, startDate
		}

		borrow.StartDate = startDate
		borrow.EndDate = endDate
		r.db[borrowID] = borrow
		return borrow, nil
	}

	return nil, utils.ErrBorrowedBookNotFound
}
