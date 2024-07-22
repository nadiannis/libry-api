package repository

import "github.com/nadiannis/libry-api/internal/domain"

type UserRepository struct {
	DB map[string]*domain.User
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
