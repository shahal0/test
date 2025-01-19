package usecase

import (
	"nproject/domain"
	"nproject/repository"
)

type UserUseCase interface {
	CreateUser(user domain.User) error
	GetUserByID(id int) (domain.User, error)
}

type UserUseCaseImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repo: repo}
}

func (uc *UserUseCaseImpl) CreateUser(user domain.User) error {
	return uc.repo.Save(user)
}

func (uc *UserUseCaseImpl) GetUserByID(id int) (domain.User, error) {
	return uc.repo.FindByID(id)
}
