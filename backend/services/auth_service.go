// services/auth_service.go
package services

import (
	"github.com/jimsyyap/auctions/backend/repositories"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}
