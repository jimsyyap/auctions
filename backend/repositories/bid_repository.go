// repositories/bid_repository.go
package repositories

import (
	"github.com/jimsyyap/auctions/backend/database"
	"github.com/jimsyyap/auctions/backend/models"
	"gorm.io/gorm"
)

type BidRepository struct {
	db *gorm.DB
}

func NewBidRepository() *BidRepository {
	return &BidRepository{
		db: database.DB,
	}
}

func (r *BidRepository) Create(bid *models.Bid) error {
	return r.db.Create(bid).Error
}

func (r *BidRepository) FindByID(id uint) (*models.Bid, error) {
	var bid models.Bid
	err := r.db.Preload("User").Preload("Listing").First(&bid, id).Error
	return &bid, err
}

func (r *BidRepository) FindByListing(listingID uint) ([]models.Bid, error) {
	var bids []models.Bid
	err := r.db.Where("listing_id = ?", listingID).
		Preload("User").
		Order("amount DESC").
		Find(&bids).Error
	return bids, err
}

func (r *BidRepository) GetHighestBid(listingID uint) (*models.Bid, error) {
	var bid models.Bid
	err := r.db.Where("listing_id = ?", listingID).
		Order("amount DESC").
		First(&bid).Error
	return &bid, err
}
