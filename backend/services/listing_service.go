// services/listing_service.go
package services

import (
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
