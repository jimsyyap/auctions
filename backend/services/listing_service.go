// services/listing_service.go
package services

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/jimsyyap/auctions/backend/models"
	"github.com/jimsyyap/auctions/backend/repositories"
)

type ListingService struct {
	listingRepo  *repositories.ListingRepository
	categoryRepo *repositories.CategoryRepository
}

func NewListingService(listingRepo *repositories.ListingRepository, categoryRepo *repositories.CategoryRepository) *ListingService {
	return &ListingService{
		listingRepo:  listingRepo,
		categoryRepo: categoryRepo,
	}
}

// ListingRequest represents data for creating/updating a listing
type ListingRequest struct {
	Title        string    `json:"title" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	StartPrice   float64   `json:"start_price" binding:"required"`
	ReservePrice float64   `json:"reserve_price"`
	BuyNowPrice  float64   `json:"buy_now_price"`
	Duration     int       `json:"duration" binding:"required"` // Duration in days
	CategoryIDs  []uint    `json:"category_ids" binding:"required"`
}

// GetListings retrieves listings with pagination
func (s *ListingService) GetListings(page, limit int) ([]models.Listing, int64, error) {
	return s.listingRepo.FindAll(page, limit)
}

// GetListing retrieves a listing by ID
func (s *ListingService) GetListing(id uint) (*models.Listing, error) {
	return s.listingRepo.FindByID(id)
}

// CreateListing creates a new listing
func (s *ListingService) CreateListing(userID uint, req *ListingRequest) (*models.Listing, error) {
	// Validate listing data
	if req.StartPrice <= 0 {
		return nil, errors.New("start price must be greater than zero")
	}

	if req.ReservePrice > 0 && req.ReservePrice < req.StartPrice {
		return nil, errors.New("reserve price must be greater than or equal to start price")
	}

	if req.BuyNowPrice > 0 && req.BuyNowPrice < req.ReservePrice {
		return nil, errors.New("buy now price must be greater than or equal to reserve price")
	}

	if req.Duration < 1 || req.Duration > 14 {
		return nil, errors.New("duration must be between 1 and 14 days")
	}

	// Fetch categories
	var categories []models.Category
	for _, categoryID := range req.CategoryIDs {
		category, err := s.categoryRepo.FindByID(categoryID)
		if err != nil {
			return nil, errors.New("invalid category ID")
		}
		categories = append(categories, *category)
	}

	// Create listing
	listing := &models.Listing{
		Title:        req.Title,
		Description:  req.Description,
		StartPrice:   req.StartPrice,
		ReservePrice: req.ReservePrice,
		BuyNowPrice:  req.BuyNowPrice,
		Status:       "active",
		EndTime:      time.Now().Add(time.Duration(req.Duration) * 24 * time.Hour),
		UserID:       userID,
		Categories:   categories,
	}

	if err := s.listingRepo.Create(listing); err != nil {
		return nil, err
	}

	return listing, nil
}

// UpdateListing updates an existing listing
func (s *ListingService) UpdateListing(id, userID uint, req *ListingRequest) (*models.Listing, error) {
	listing, err := s.listingRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("listing not found")
	}

	// Verify ownership
	if listing.UserID != userID {
		return nil, errors.New("you do not have permission to update this listing")
	}

	// Verify listing is still active and has no bids
	if listing.Status != "active" {
		return nil, errors.New("only active listings can be updated")
	}

	// Check if there are any bids
	if len(listing.Bids) > 0 {
		return nil, errors.New("listings with bids cannot be updated")
	}

	// Validate listing data
	if req.StartPrice <= 0 {
		return nil, errors.New("start price must be greater than zero")
	}

	if req.ReservePrice > 0 && req.ReservePrice < req.StartPrice {
		return nil, errors.New("reserve price must be greater than or equal to start price")
	}

	if req.BuyNowPrice > 0 && req.BuyNowPrice < req.ReservePrice {
		return nil, errors.New("buy now price must be greater than or equal to reserve price")
	}

	// Update fields
	listing.Title = req.Title
	listing.Description = req.Description
	listing.StartPrice = req.StartPrice
	listing.ReservePrice = req.ReservePrice
	listing.BuyNowPrice = req.BuyNowPrice

	// Only allow duration update if the listing has no bids
	if req.Duration >= 1 && req.Duration <= 14 {
		// Calculate new end time based on creation time
		createdAt := listing.CreatedAt
		listing.EndTime = createdAt.Add(time.Duration(req.Duration) * 24 * time.Hour)
	}

	// Update categories if provided
	if len(req.CategoryIDs) > 0 {
		var categories []models.Category
		for _, categoryID := range req.CategoryIDs {
			category, err := s.categoryRepo.FindByID(categoryID)
			if err != nil {
				return nil, errors.New("invalid category ID")
			}
			categories = append(categories, *category)
		}
		listing.Categories = categories
	}

	if err := s.listingRepo.Update(listing); err != nil {
		return nil, err
	}

	return listing, nil
}

// DeleteListing deletes a listing
func (s *ListingService) DeleteListing(id, userID uint) error {
	listing, err := s.listingRepo.FindByID(id)
	if err != nil {
		return errors.New("listing not found")
	}

	// Verify ownership
	if listing.UserID != userID {
		return errors.New("you do not have permission to delete this listing")
	}

	// Verify listing has no bids
	if len(listing.Bids) > 0 {
		return errors.New("listings with bids cannot be deleted")
	}

	// Delete any associated images from the filesystem
	for _, image := range listing.Images {
		// Get file path
		imgPath := filepath.Join("uploads", "listings", strconv.FormatUint(uint64(listing.ID), 10), image.Filename)
		
		// Attempt to remove the file
		err := os.Remove(imgPath)
		if err != nil && !os.IsNotExist(err) {
			// Log error but continue with deletion process
			// You might want to add proper logging here
			// log.Printf("Error deleting image %s: %v", imgPath, err)
		}
	}

	// Delete the listing from the database
	return s.listingRepo.Delete(id)
}
