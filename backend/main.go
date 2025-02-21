package main

import (
	"github.com/gin-gonic/gin"
	"your-project-name/backend/routes"
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
