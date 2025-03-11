// services/user_service.go
package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

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

// UserProfile represents a user's public profile
type UserProfile struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	Rating      float64 `json:"rating"`
}

// UpdateProfileRequest represents the data for updating a user profile
type UpdateProfileRequest struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	// Don't expose the password hash
	user.Password = ""
	return user, nil
}

// GetUserProfile creates a safe user profile from a user model
func (s *UserService) GetUserProfile(id uint) (*UserProfile, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// TODO: Calculate user rating
	// This would typically be done by averaging all ratings received by the user
	rating := 0.0

	return &UserProfile{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Rating:      rating,
	}, nil
}

// UpdateProfile updates a user's profile information
func (s *UserService) UpdateProfile(userID uint, req *UpdateProfileRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	// If changing email, check if it's already used
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := s.userRepo.FindByEmail(req.Email)
		if err == nil && existingUser.ID > 0 && existingUser.ID != userID {
			return errors.New("email already in use")
		}
		user.Email = req.Email
	}

	// Update user fields if provided
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}
	if req.Address != "" {
		user.Address = req.Address
	}

	// Change password if requested
	if req.Password != "" && req.NewPassword != "" {
		// Verify old password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			return errors.New("current password is incorrect")
		}

		// Check password requirements
		authService := AuthService{}
		if err := authService.validatePassword(req.NewPassword); err != nil {
			return err
		}

		// Hash new password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return s.userRepo.Update(user)
}

// GetUserListings gets all listings by a specific user
func (s *UserService) GetUserListings(userID uint, page, limit int) ([]models.Listing, int64, error) {
	return s.userRepo.GetUserListings(userID, page, limit)
}

// GetUserBids gets all bids placed by a specific user
func (s *UserService) GetUserBids(userID uint, page, limit int) ([]models.Bid, int64, error) {
	return s.userRepo.GetUserBids(userID, page, limit)
}
