// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/auctions/backend/handlers"
	"github.com/jimsyyap/auctions/backend/middlewares"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, handlers *handlers.Handlers) {
	// API group
	api := router.Group("/api")
	
	// Setup route groups
	setupAuthRoutes(api, handlers.AuthHandler)
	setupUserRoutes(api, handlers.UserHandler)
	setupListingRoutes(api, handlers.ListingHandler)
	setupCategoryRoutes(api, handlers.CategoryHandler)
	setupBidRoutes(api, handlers.BidHandler)
}

// Individual route group setups...
