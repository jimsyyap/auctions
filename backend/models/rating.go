// models/rating.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Score        int       `gorm:"not null;check:score >= 1 AND score <= 5"`
	Comment      string
	CreatedAt    time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	
	// Who gave the rating
	RaterUserID  uint
	RaterUser    User      `gorm:"foreignKey:RaterUserID"`
	
	// Who received the rating
	RatedUserID  uint
	RatedUser    User      `gorm:"foreignKey:RatedUserID"`
	
	// Related to which listing
	ListingID    uint
	Listing      Listing   `gorm:"foreignKey:ListingID"`
}

// BeforeCreate validates the rating
func (r *Rating) BeforeCreate(tx *gorm.DB) error {
	// Ensure score is between 1 and 5
	if r.Score < 1 {
		r.Score = 1
	} else if r.Score > 5 {
		r.Score = 5
	}
	
	// Set creation time
	r.CreatedAt = time.Now()
	
	return nil
}
