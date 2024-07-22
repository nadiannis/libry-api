package repository

type Repositories struct {
	Books   IBookRepository
	Users   IUserRepository
	Borrows IBorrowRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Books:   NewBookRepository(),
		Users:   NewUserRepository(),
		Borrows: NewBorrowRepository(),
	}
}
