// repositories/category_repository.go
package repositories

import (
	"github.com/jimsyyap/auctions/backend/database"
	"github.com/jimsyyap/auctions/backend/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		db: database.DB,
	}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Children").First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Children").Where("parent_id IS NULL").Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetListingsByCategory(categoryID uint, page, limit int) ([]models.Listing, int64, error) {
	var listings []models.Listing
	var count int64
	
	offset := (page - 1) * limit
	
	// Query to get listings in this category or its children
	query := r.db.Joins("JOIN listing_categories ON listing_categories.listing_id = listings.id").
		Where("listing_categories.category_id = ?", categoryID)
	
	err := query.Model(&models.Listing{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = query.Preload("User").
		Preload("Categories").
		Preload("Images").
		Offset(offset).Limit(limit).
		Find(&listings).Error
	
	return listings, count, err
}
