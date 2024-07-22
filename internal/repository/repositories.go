package repository

import "github.com/nadiannis/libry-api/internal/domain"

type Repositories struct {
	Books   BookRepository
	Users   UserRepository
	Borrows BorrowRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Books: BookRepository{
			DB: make(map[string]*domain.Book),
		},
		Users: UserRepository{
			DB: make(map[string]*domain.User),
		},
		Borrows: BorrowRepository{
			DB: make(map[string]*domain.Borrow),
		},
	}
}
