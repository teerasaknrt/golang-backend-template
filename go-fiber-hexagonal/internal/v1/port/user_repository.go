package port

import "go-fiber-hexagonal/internal/v1/domain"

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
}
