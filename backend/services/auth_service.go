// services/auth_service.go
package services

import (
	"errors"
	"regexp"
	"time"

	"github.com/jimsyyap/auctions/backend/middlewares"
	"github.com/jimsyyap/auctions/backend/models"
	"github.com/jimsyyap/auctions/backend/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

// LoginRequest represents login form data
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents registration form data
type RegisterRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

// AuthResponse represents the response after successful authentication
type AuthResponse struct {
	Token     string      `json:"token"`
	ExpiresAt time.Time   `json:"expires_at"`
	User      *models.User `json:"user"`
}

// Register creates a new user account
func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
	// Check if username already exists
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	existingUser, err = s.userRepo.FindByEmail(req.Email)
	if err == nil && existingUser.ID > 0 {
		return nil, errors.New("email already exists")
	}

	// Validate password
	if err := s.validatePassword(req.Password); err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    string(hashedPassword),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		IsAdmin:     false,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := middlewares.GenerateToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		User:      user,
	}, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(req *LoginRequest) (*AuthResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := middlewares.GenerateToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	// Sanitize user data before returning
	user.Password = ""

	return &AuthResponse{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		User:      user,
	}, nil
}

// RefreshToken generates a new token for a valid user
func (s *AuthService) RefreshToken(userID uint) (*AuthResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Generate new JWT token
	token, err := middlewares.GenerateToken(user.ID, user.Username, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	// Sanitize user data before returning
	user.Password = ""

	return &AuthResponse{
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		User:      user,
	}, nil
}

// validatePassword checks password against security requirements
func (s *AuthService) validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Check for at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for at least one digit
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must contain at least one digit")
	}

	// Check for at least one special character
	if !regexp.MustCompile(`[!@#$%^&*]`).MatchString(password) {
		return errors.New("password must contain at least one special character (!@#$%^&*)")
	}

	return nil
}
