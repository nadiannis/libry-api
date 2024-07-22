package handler

import "github.com/nadiannis/libry-api/internal/usecase"

type Handlers struct {
	Books BookHandler
	Users UserHandler
}

func NewHandlers() Handlers {
	return Handlers{
		Books: NewBookHandler(usecase.NewUsecases().Books),
		Users: NewUsersHandler(usecase.NewUsecases().Users),
	}
}

func NewBookHandler(usecase usecase.BookUsecase) BookHandler {
	return BookHandler{
		usecase: usecase,
	}
}

func NewUsersHandler(usecase usecase.UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}
