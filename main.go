package main

import (
	"net/http"

	"github.com/thiagocprado/golang-hexagonal-architecture/internal/adapter/handler"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/adapter/repository"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/domain/usecase"
)

func main() {
	r := repository.NewMemoryUserRepository()
	u := usecase.NewUserUseCase(r)
	h := handler.NewUserHandler(u)
	router := handler.NewRouter(h)

	http.ListenAndServe(":8080", router)
}
