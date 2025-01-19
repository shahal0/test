package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "nproject/domain"
    "nproject/handlers"
    "testing"

    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

// MockUsecase defines a mock implementation of the usecase interface.
type MockUsecase struct {
    mock.Mock
}

func (m *MockUsecase) CreateUser(user domain.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUsecase) GetUserByID(id int) (domain.User, error) {
    args := m.Called(id)
    return args.Get(0).(domain.User), args.Error(1)
}

func TestCreateUserHandler(t *testing.T) {
    mockUsecase := new(MockUsecase)
    handler := handlers.NewUserHandler(mockUsecase)

    user := domain.User{Name: "John Doe", Email: "john@example.com"}
    mockUsecase.On("CreateUser", user).Return(nil)

    // Create a request and recorder
    body, _ := json.Marshal(user)
    req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    rec := httptest.NewRecorder()

    // Call the handler
    handler.CreateUser(rec, req)

    // Assertions
    require.Equal(t, http.StatusCreated, rec.Code)
    mockUsecase.AssertCalled(t, "CreateUser", user)
}
