package usecase

import "github.com/nadiannis/libry-api/internal/repository"

type Usecases struct {
	Books   BookUsecase
	Users   UserUsecase
	Borrows BorrowUsecase
}

func NewUsecases() Usecases {
	return Usecases{
		Books:   NewBookUsecase(repository.NewRepositories().Books),
		Users:   NewUserUsecase(repository.NewRepositories().Users),
		Borrows: NewBorrowUsecase(repository.NewRepositories().Borrows),
	}
}
