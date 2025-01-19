package main

import (
	"log"
	"net/http"
	"nproject/handlers"
	"nproject/repository"
	"nproject/usecase"
)

func main() {
	// Initialize repository, use case, and handler
	repo := repository.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	handler := handlers.NewUserHandler(uc)

	// Register routes
	handler.RegisterRoutes()

	// Start the HTTP server
	log.Println("Server running on http://0.0.0.0:8000")
	err := http.ListenAndServe(":8000", nil) // Listen on all interfaces on port 8000
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
