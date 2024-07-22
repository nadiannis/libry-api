package usecase

import "github.com/nadiannis/libry-api/internal/repository"

type Usecases struct {
	Books BookUsecase
}

func NewUsecases() Usecases {
	return Usecases{
		Books: NewBookUsecase(repository.NewRepositories().Books),
	}
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return BookUsecase{
		repository: repository,
	}
}
