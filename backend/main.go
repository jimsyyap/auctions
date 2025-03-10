// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/auctions/backend/config"
	"github.com/jimsyyap/auctions/backend/database"
	"github.com/jimsyyap/auctions/backend/handlers"
	"github.com/jimsyyap/auctions/backend/middlewares"
	"github.com/jimsyyap/auctions/backend/repositories"
	"github.com/jimsyyap/auctions/backend/routes"
	"github.com/jimsyyap/auctions/backend/services"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	database.Connect()

	// Run migrations
	database.Migrate()

	// Initialize repositories
	userRepo := repositories.NewUserRepository()
	listingRepo := repositories.NewListingRepository()
	bidRepo := repositories.NewBidRepository()
	categoryRepo := repositories.NewCategoryRepository()

	// Initialize services
	userService := services.NewUserService(userRepo)
	listingService := services.NewListingService(listingRepo, categoryRepo)
	bidService := services.NewBidService(bidRepo, listingRepo, userRepo)
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	listingHandler := handlers.NewListingHandler(listingService)
	bidHandler := handlers.NewBidHandler(bidService)
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Apply global middlewares
	router.Use(middlewares.Logger())
	router.Use(middlewares.Recovery())

	// Set up API routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// User routes
		users := api.Group("/users")
		users.Use(middlewares.Auth()) // Require authentication
		{
			users.GET("/me", userHandler.GetProfile)
			users.PUT("/me", userHandler.UpdateProfile)
			users.GET("/:id", userHandler.GetUser)
			users.GET("/:id/listings", userHandler.GetUserListings)
			users.GET("/:id/bids", userHandler.GetUserBids)
		}

		// Listing routes
		listings := api.Group("/listings")
		{
			listings.GET("", listingHandler.GetListings)
			listings.GET("/:id", listingHandler.GetListing)
			
			// Protected routes
			authenticated := listings.Group("")
			authenticated.Use(middlewares.Auth())
			{
				authenticated.POST("", listingHandler.CreateListing)
				authenticated.PUT("/:id", listingHandler.UpdateListing)
				authenticated.DELETE("/:id", listingHandler.DeleteListing)
				authenticated.POST("/:id/bids", bidHandler.PlaceBid)
			}
		}

		// Category routes
		categories := api.Group("/categories")
		{
			categories.GET("", listingHandler.GetCategories)
			categories.GET("/:id/listings", listingHandler.GetListingsByCategory)
		}
	}

	// Add health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
			"time":   time.Now(),
		})
	})

	// Get server configuration
	serverConfig := config.GetServerConfig()
	serverAddr := fmt.Sprintf(":%s", serverConfig.Port)

	// Start the server
	log.Printf("Server starting on port %s", serverConfig.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
