package repository

import (
	"errors"

	"github.com/thiagocprado/golang-hexagonal-architecture/internal/core/entity"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/port/output"
)

type MemoryUserRepository struct {
	data map[string]entity.User
}

func NewMemoryUserRepository() output.UserRepository {
	return &MemoryUserRepository{data: make(map[string]entity.User)}
}

func (r *MemoryUserRepository) Save(user entity.User) error {
	r.data[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) FindByID(id string) (entity.User, error) {
	user, exists := r.data[id]
	if !exists {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}
