package usecase

import (
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/repository"
)

type BorrowUsecase struct {
	repository repository.BorrowRepository
}

func NewBorrowUsecase(repository repository.BorrowRepository) BorrowUsecase {
	return BorrowUsecase{
		repository: repository,
	}
}

func (u *BorrowUsecase) GetAll() []*domain.Borrow {
	return u.repository.GetAll()
}
