// handlers/user_handler.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/auctions/backend/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile retrieves the logged-in user's profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	
	profile, err := h.userService.GetUserProfile(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}
	
	c.JSON(http.StatusOK, profile)
}

// UpdateProfile updates the logged-in user's profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	
	var req services.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.userService.UpdateProfile(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// GetUser retrieves a public user profile by ID
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	profile, err := h.userService.GetUserProfile(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	c.JSON(http.StatusOK, profile)
}

// GetUserListings gets all listings by a specific user
func (h *UserHandler) GetUserListings(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	listings, total, err := h.userService.GetUserListings(uint(id), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user listings"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"listings": listings,
		"pagination": gin.H{
			"total": total,
			"page": page,
			"limit": limit,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetUserBids gets all bids placed by a specific user
func (h *UserHandler) GetUserBids(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	bids, total, err := h.userService.GetUserBids(uint(id), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user bids"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"bids": bids,
		"pagination": gin.H{
			"total": total,
			"page": page,
			"limit": limit,
			"pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}
