package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/jimsyyap/auctions/backend/utils"
)

func main() {
    // Connect to the database
    db := utils.ConnectDB()
    defer db.Close()

    // Initialize Gin
    r := gin.Default()

    // Enable CORS middleware
    r.Use(cors.Default())

    // If you want to customize the CORS configuration (e.g., to allow only specific origins or methods), you can replace cors.Default() with a custom configuration.
    // r.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"http://localhost:3000"}, // Allow only your frontend origin
    //     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    //     AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     MaxAge:           12 * time.Hour,
    // }))

    // Example route
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Auction Website!",
        })
    })

    // Start the server
    r.Run() // Listen and serve on 0.0.0.0:8080
}
