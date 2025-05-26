package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/thiagocprado/golang-hexagonal-architecture/internal/domain/entity"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/port/input"
)

type UserHandler struct {
	useCase input.UserUseCase
}

func NewUserHandler(useCase input.UserUseCase) *UserHandler {
	return &UserHandler{useCase}
}

func (h *UserHandler) Save(w http.ResponseWriter, r *http.Request) {
	var body entity.User

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		fmt.Println("Error decoding request body:", err.Error())
		return
	}

	user, err := h.useCase.Save(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]

	user, err := h.useCase.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
