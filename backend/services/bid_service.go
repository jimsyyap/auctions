// services/bid_service.go
package services

import (
	"github.com/jimsyyap/auctions/backend/repositories"
)

type BidService struct {
	bidRepo     *repositories.BidRepository
	listingRepo *repositories.ListingRepository
	userRepo    *repositories.UserRepository
}

func NewBidService(bidRepo *repositories.BidRepository, listingRepo *repositories.ListingRepository, userRepo *repositories.UserRepository) *BidService {
	return &BidService{
		bidRepo:     bidRepo,
		listingRepo: listingRepo,
		userRepo:    userRepo,
	}
}
