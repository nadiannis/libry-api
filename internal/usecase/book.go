package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/repository"
)

type BookUsecase struct {
	repository repository.BookRepository
}

func (u *BookUsecase) GetAll() []*domain.Book {
	return u.repository.GetAll()
}

func (u *BookUsecase) Add(body *request.BookRequest) *domain.Book {
	book := &domain.Book{
		ID:     uuid.NewString(),
		Title:  body.Title,
		Author: body.Author,
	}

	return u.repository.Add(book)
}
