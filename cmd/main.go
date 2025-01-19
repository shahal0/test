package main

import (
	"nproject/handlers"
	"nproject/repository"
	"nproject/usecase"
)

func main() {
	repo := repository.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	handler := handlers.NewUserHandler(uc)

	handler.RegisterRoutes()
}
