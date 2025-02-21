package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/auctions/backend/handlers"
	"github.com/jimsyyap/auctions/backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	// Protected routes
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handlers.Profile) // Example protected route
	}

	// Admin-only routes
	admin := auth.Group("/admin")
	admin.Use(middleware.RBACMiddleware("Admin"))
	{
		admin.GET("/users", handlers.ListUsers) // Example admin route
	}
}
