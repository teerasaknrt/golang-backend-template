package port

import "go-fiber-hexagonal/internal/v1/domain"

type UserUseCase interface {
	GetUserByID(id uint) (*domain.User, error)
}
