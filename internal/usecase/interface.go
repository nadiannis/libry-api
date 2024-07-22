package usecase

import (
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/domain/request"
)

type IUserUsecase interface {
	GetAll() []*domain.User
	Add(input *request.UserRequest) *domain.User
}

type IBookUsecase interface {
	GetAll() []*domain.Book
	Add(body *request.BookRequest) *domain.Book
}

type IBorrowUsecase interface {
	GetAll() []*domain.Borrow
	Borrow(body *request.BorrowRequest) (*domain.Borrow, error)
	Return(body *request.BorrowRequest) (*domain.Borrow, error)
}
