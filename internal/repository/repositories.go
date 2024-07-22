package repository

import "github.com/nadiannis/libry-api/internal/domain"

type Repositories struct {
	Books BookRepository
}

func NewRepositories() Repositories {
	return Repositories{
		Books: BookRepository{
			DB: make(map[string]*domain.Book),
		},
	}
}
