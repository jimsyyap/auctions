// models/bid.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Bid struct {
	gorm.Model
	Amount      float64   `gorm:"not null"`
	PlacedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	
	// Relationships
	UserID      uint
	User        User      `gorm:"foreignKey:UserID"`
	ListingID   uint
	Listing     Listing   `gorm:"foreignKey:ListingID"`
}

// BeforeCreate sets the PlacedAt field to the current time
func (b *Bid) BeforeCreate(tx *gorm.DB) error {
	b.PlacedAt = time.Now()
	return nil
}
