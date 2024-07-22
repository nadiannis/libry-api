package handler

import "github.com/nadiannis/libry-api/internal/usecase"

type Handlers struct {
	Books   IBookHandler
	Users   IUserHandler
	Borrows IBorrowHandler
}

func NewHandlers(usecases usecase.Usecases) Handlers {
	return Handlers{
		Books:   NewBookHandler(usecases.Books),
		Users:   NewUserHandler(usecases.Users),
		Borrows: NewBorrowHandler(usecases.Borrows),
	}
}
