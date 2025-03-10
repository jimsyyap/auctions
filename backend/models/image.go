// models/image.go
package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	URL         string  `gorm:"not null"`
	Caption     string
	IsPrimary   bool    `gorm:"default:false"`
	DisplayOrder int    `gorm:"default:0"`
	
	// Relationships
	ListingID   uint
	Listing     Listing `gorm:"foreignKey:ListingID"`
}

// BeforeCreate ensures only one image is primary
func (i *Image) BeforeCreate(tx *gorm.DB) error {
	// If this image is set as primary, un-set any existing primary for this listing
	if i.IsPrimary {
		tx.Model(&Image{}).Where("listing_id = ? AND is_primary = ?", i.ListingID, true).Update("is_primary", false)
	}
	return nil
}
