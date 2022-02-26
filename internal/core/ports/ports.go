package ports

import (
	"hexagonal-boilerplate/internal/core/domain"
)

type UserRepository interface {
	Save(user domain.User) error
	GetUser(email string) (domain.User, error)
}
