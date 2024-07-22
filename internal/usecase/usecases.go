package usecase

import "github.com/nadiannis/libry-api/internal/repository"

type Usecases struct {
	Books BookUsecase
	Users UserUsecase
}

func NewUsecases() Usecases {
	return Usecases{
		Books: NewBookUsecase(repository.NewRepositories().Books),
		Users: NewUserUsecase(repository.NewRepositories().Users),
	}
}

func NewBookUsecase(repository repository.BookRepository) BookUsecase {
	return BookUsecase{
		repository: repository,
	}
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repository,
	}
}
