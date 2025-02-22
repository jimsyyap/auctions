package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/auctions/backend/routes"
    "github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Set up routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
