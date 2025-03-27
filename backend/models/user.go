// models/user.go
package models

import (
	_ "time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"uniqueIndex;not null"`
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
	IsAdmin     bool `gorm:"default:false"`
	
	// Relationships
	Listings    []Listing `gorm:"foreignKey:UserID"`
	Bids        []Bid     `gorm:"foreignKey:UserID"`
	Ratings     []Rating  `gorm:"foreignKey:RatedUserID"`
	GivenRatings []Rating `gorm:"foreignKey:RaterUserID"`
}
