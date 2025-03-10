// repositories/user_repository.go
package repositories

import (
	"github.com/jimsyyap/auctions/backend/database"
	"github.com/jimsyyap/auctions/backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetUserListings(userID uint, page, limit int) ([]models.Listing, int64, error) {
	var listings []models.Listing
	var count int64
	
	offset := (page - 1) * limit
	
	err := r.db.Model(&models.Listing{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = r.db.Where("user_id = ?", userID).
		Preload("Categories").
		Preload("Images").
		Offset(offset).Limit(limit).
		Find(&listings).Error
	
	return listings, count, err
}

func (r *UserRepository) GetUserBids(userID uint, page, limit int) ([]models.Bid, int64, error) {
	var bids []models.Bid
	var count int64
	
	offset := (page - 1) * limit
	
	err := r.db.Model(&models.Bid{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = r.db.Where("user_id = ?", userID).
		Preload("Listing").
		Preload("Listing.Images").
		Offset(offset).Limit(limit).
		Find(&bids).Error
	
	return bids, count, err
}
