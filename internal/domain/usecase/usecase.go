package usecase

import (
	"github.com/google/uuid"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/domain/entity"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/port/input"
	"github.com/thiagocprado/golang-hexagonal-architecture/internal/port/output"
)

type UserUseCase struct {
	repo output.UserRepository
}

func NewUserUseCase(repo output.UserRepository) input.UserUseCase {
	return &UserUseCase{repo}
}

func (s *UserUseCase) Save(body entity.User) (entity.User, error) {
	user, err := entity.NewUser(uuid.New().String(), body.Name, body.Age)
	if err != nil {
		return entity.User{}, err
	}

	if err := s.repo.Save(user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *UserUseCase) FindByID(id string) (entity.User, error) {
	return s.repo.FindByID(id)
}
