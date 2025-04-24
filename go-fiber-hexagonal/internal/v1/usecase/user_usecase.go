package usecase

import (
	"go-fiber-hexagonal/internal/v1/domain"
	"go-fiber-hexagonal/internal/v1/port"
)

type userUseCase struct {
	repo port.UserRepository
}

func NewUserUseCase(repo port.UserRepository) port.UserUseCase {
	return &userUseCase{repo}
}

func (u *userUseCase) GetUserByID(id uint) (*domain.User, error) {
	return u.repo.FindByID(id)
}
