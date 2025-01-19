package usecase_test

import (
	"nproject/domain"
	"nproject/repository"
	"nproject/usecase"
	"testing"
)

func TestCreateUser(t *testing.T) {
	mockRepo := &repository.MockRepository{
		SaveFunc: func(user domain.User) error {
			if user.Name == "" || user.Email == "" {
				t.Errorf("Save called with invalid user: %+v", user)
			}
			return nil
		},
	}

	uc := usecase.NewUserUseCase(mockRepo)

	user := domain.User{Name: "Test User", Email: "test@example.com"}
	err := uc.CreateUser(user)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestGetUserByID(t *testing.T) {
	mockRepo := &repository.MockRepository{
		FindByIDFunc: func(id int) (domain.User, error) {
			if id == 1 {
				return domain.User{ID: 1, Name: "Test User", Email: "test@example.com"}, nil
			}
			return domain.User{}, nil
		},
	}

	uc := usecase.NewUserUseCase(mockRepo)

	user, err := uc.GetUserByID(1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if user.ID != 1 || user.Name != "Test User" || user.Email != "test@example.com" {
		t.Errorf("unexpected user: %+v", user)
	}
}
