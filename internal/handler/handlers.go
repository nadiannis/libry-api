package handler

import "github.com/nadiannis/libry-api/internal/usecase"

type Handlers struct {
	Books BookHandler
}

func NewHandlers() Handlers {
	return Handlers{
		Books: NewBookHandler(usecase.NewUsecases().Books),
	}
}

func NewBookHandler(usecase usecase.BookUsecase) BookHandler {
	return BookHandler{
		usecase: usecase,
	}
}
