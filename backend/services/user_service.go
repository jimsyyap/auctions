// services/user_service.go
package services

import (
	"github.com/jimsyyap/auctions/backend/models"
	"github.com/jimsyyap/auctions/backend/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}
