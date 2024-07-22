package usecase

import (
	"github.com/google/uuid"
	"github.com/nadiannis/libry-api/internal/domain"
	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/repository"
)

type UserUsecase struct {
	repository repository.IUserRepository
}

func NewUserUsecase(repository repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (u *UserUsecase) GetAll() []*domain.User {
	return u.repository.GetAll()
}

func (u *UserUsecase) Add(input *request.UserRequest) *domain.User {
	user := &domain.User{
		ID:       uuid.NewString(),
		Username: input.Username,
		Books:    make([]*domain.Book, 0),
	}

	return u.repository.Add(user)
}
