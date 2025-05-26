package output

import "github.com/thiagocprado/golang-hexagonal-architecture/internal/core/entity"

type UserRepository interface {
	Save(user entity.User) error
	FindByID(id string) (entity.User, error)
}
