// repositories/listing_repository.go
package repositories

import (
    "github.com/jimsyyap/auctions/backend/database"
    "github.com/jimsyyap/auctions/backend/models"
    "gorm.io/gorm"
)

type ListingRepository struct {
    db *gorm.DB
}

func NewListingRepository() *ListingRepository {
    return &ListingRepository{
        db: database.DB,
    }
}

func (r *ListingRepository) Create(listing *models.Listing) error {
    return r.db.Create(listing).Error
}

func (r *ListingRepository) FindByID(id uint) (*models.Listing, error) {
    var listing models.Listing
    err := r.db.Preload("User").Preload("Categories").Preload("Images").First(&listing, id).Error
    return &listing, err
}

func (r *ListingRepository) FindAll(page, limit int) ([]models.Listing, int64, error) {
    var listings []models.Listing
    var count int64
    
    offset := (page - 1) * limit
    
    err := r.db.Model(&models.Listing{}).Count(&count).Error
    if err != nil {
        return nil, 0, err
    }
    
    err = r.db.Preload("User").Preload("Categories").Preload("Images").
           Offset(offset).Limit(limit).Find(&listings).Error
    
    return listings, count, err
}

func (r *ListingRepository) Update(listing *models.Listing) error {
    return r.db.Save(listing).Error
}

func (r *ListingRepository) Delete(id uint) error {
    return r.db.Delete(&models.Listing{}, id).Error
}

// Add more query methods as needed
