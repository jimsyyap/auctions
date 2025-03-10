// database/migration.go
package database

import (
    "log"

    "github.com/jimsyyap/auctions/backend/models"
)

func Migrate() {
    log.Println("Running database migrations...")
    
    // Add all your models here
    err := DB.AutoMigrate(
        &models.User{},
        &models.Listing{},
        &models.Bid{},
        &models.Category{},
        &models.Image{},
        &models.Rating{},
    )
    
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    
    log.Println("Database migration completed")
}
