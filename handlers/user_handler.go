package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nproject/domain"
	"nproject/usecase"
	"strconv"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) RegisterRoutes() {
	http.HandleFunc("/user", h.CreateUser)
	http.HandleFunc("/user/get", h.GetUserByID)
	log.Println("Server running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.usecase.CreateUser(user); err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
