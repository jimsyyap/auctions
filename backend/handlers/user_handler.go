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

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	// Implementation
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserListings(c *gin.Context) {
	// Implementation
	c.JSON(http.StatusOK, gin.H{"listings": []string{}})
}

func (h *UserHandler) GetUserBids(c *gin.Context) {
	// Implementation
	c.JSON(http.StatusOK, gin.H{"bids": []string{}})
}
