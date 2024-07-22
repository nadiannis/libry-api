package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/repository"
)

type BorrowUsecase struct {
	borrowRepository repository.IBorrowRepository
	bookRepository   repository.IBookRepository
	userRepository   repository.IUserRepository
}

func NewBorrowUsecase(
	borrowRepository repository.IBorrowRepository,
	bookRepository repository.IBookRepository,
	userRepository repository.IUserRepository,
) IBorrowUsecase {
	return &BorrowUsecase{
		borrowRepository: borrowRepository,
		bookRepository:   bookRepository,
		userRepository:   userRepository,
	}
}

func (u *BorrowUsecase) GetAll() []*domain.Borrow {
	return u.borrowRepository.GetAll()
}

func (u *BorrowUsecase) Borrow(body *request.BorrowRequest) (*domain.Borrow, error) {
	user, err := u.userRepository.GetByID(body.UserID)
	if err != nil {
		return nil, err
	}

	book, err := u.bookRepository.GetByID(body.BookID)
	if err != nil {
		return nil, err
	}

	borrow := &domain.Borrow{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		BookID:    book.ID,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 7),
		Status:    domain.StatusBorrowed,
	}

	borrowedBook, err := u.borrowRepository.Borrow(borrow)
	if err != nil {
		return nil, err
	}

	u.userRepository.AddBook(user.ID, book)

	return borrowedBook, nil
}
