package handler

import "github.com/nadiannis/libry-api/internal/usecase"

type Handlers struct {
	Books   BookHandler
	Users   UserHandler
	Borrows BorrowHandler
}

func NewHandlers() Handlers {
	return Handlers{
		Books:   NewBookHandler(usecase.NewUsecases().Books),
		Users:   NewUserHandler(usecase.NewUsecases().Users),
		Borrows: NewBorrowHandler(usecase.NewUsecases().Borrows),
	}
}
