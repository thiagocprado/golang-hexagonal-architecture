package input

import "github.com/thiagocprado/golang-hexagonal-architecture/internal/domain/entity"

type UserUseCase interface {
	Save(body entity.User) (entity.User, error)
	FindByID(id string) (entity.User, error)
}
