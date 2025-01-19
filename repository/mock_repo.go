package repository

import "nproject/domain"

type MockRepository struct {
	SaveFunc     func(user domain.User) error
	FindByIDFunc func(id int) (domain.User, error)
}

func (m *MockRepository) Save(user domain.User) error {
	return m.SaveFunc(user)
}

func (m *MockRepository) FindByID(id int) (domain.User, error) {
	return m.FindByIDFunc(id)
}
