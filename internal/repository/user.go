package repository

import (
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/utils"
)

type UserRepository struct {
	DB map[string]*domain.User
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		DB: make(map[string]*domain.User),
	}
}

func (r *UserRepository) GetAll() []*domain.User {
	users := make([]*domain.User, 0)
	for _, user := range r.DB {
		users = append(users, user)
	}
	return users
}

func (r *UserRepository) Add(user *domain.User) *domain.User {
	r.DB[user.ID] = user
	return user
}

func (r *UserRepository) GetByID(userID string) (*domain.User, error) {
	if user, exists := r.DB[userID]; exists {
		return user, nil
	}

	return nil, utils.ErrUserNotFound
}

func (r *UserRepository) AddBook(userID string, book *domain.Book) (*domain.Book, error) {
	if user, exists := r.DB[userID]; exists {
		user.Books = append(user.Books, book)
		return book, nil
	}

	return nil, utils.ErrUserNotFound
}

func (r *UserRepository) DeleteBookByID(userID, bookID string) error {
	user, exists := r.DB[userID]
	if !exists {
		return utils.ErrUserNotFound
	}

	for i, book := range user.Books {
		if book.ID == bookID {
			user.Books = append(user.Books[:i], user.Books[i+1:]...)
		}
	}

	return nil
}
