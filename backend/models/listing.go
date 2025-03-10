// models/listing.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Listing struct {
	gorm.Model
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"type:text"`
	StartPrice   float64   `gorm:"not null"`
	ReservePrice float64
	BuyNowPrice  float64
	Status       string    `gorm:"default:'active'"`  // active, ended, sold
	EndTime      time.Time
	
	// Categories - using a many-to-many relationship
	Categories   []Category `gorm:"many2many:listing_categories;"`
	
	// Relationships
	UserID       uint
	User         User       `gorm:"foreignKey:UserID"`
	Bids         []Bid      `gorm:"foreignKey:ListingID"`
	Images       []Image    `gorm:"foreignKey:ListingID"`
	Ratings      []Rating   `gorm:"foreignKey:ListingID"`
}

// GetCurrentPrice returns the current highest bid amount or the start price if no bids
func (l *Listing) GetCurrentPrice() float64 {
	if len(l.Bids) == 0 {
		return l.StartPrice
	}
	
	// Find the highest bid
	highestBid := l.Bids[0]
	for _, bid := range l.Bids {
		if bid.Amount > highestBid.Amount {
			highestBid = bid
		}
	}
	
	return highestBid.Amount
}

// IsReserveReached checks if the reserve price has been reached
func (l *Listing) IsReserveReached() bool {
	currentPrice := l.GetCurrentPrice()
	return currentPrice >= l.ReservePrice
}
