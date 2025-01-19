package repository

import "nproject/domain"

type UserRepository interface {
	Save(user domain.User) error
	FindByID(id int) (domain.User, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Save(user domain.User) error {
	// Simulate saving user
	return nil
}

func (repo *UserRepositoryImpl) FindByID(id int) (domain.User, error) {
	// Simulate finding user by ID
	if id == 1 {
		return domain.User{ID: 1, Name: "Test User", Email: "test@example.com"}, nil
	}
	return domain.User{}, nil
}
