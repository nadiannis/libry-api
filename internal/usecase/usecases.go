package usecase

import "github.com/nadiannis/libry-api/internal/repository"

type Usecases struct {
	Books   IBookUsecase
	Users   IUserUsecase
	Borrows IBorrowUsecase
}

func NewUsecases(repositories repository.Repositories) Usecases {
	return Usecases{
		Books: NewBookUsecase(repositories.Books),
		Users: NewUserUsecase(repositories.Users),
		Borrows: NewBorrowUsecase(
			repositories.Borrows,
			repositories.Books,
			repositories.Users,
		),
	}
}
