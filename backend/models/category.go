// models/category.go
package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	ParentID    *uint
	Parent      *Category  `gorm:"foreignKey:ParentID"`
	Children    []Category `gorm:"foreignKey:ParentID"`
	
	// Many-to-many relationship with listings
	Listings    []Listing `gorm:"many2many:listing_categories;"`
}

// IsParentCategory returns true if this category is a top-level/parent category
func (c *Category) IsParentCategory() bool {
	return c.ParentID == nil
}

// GetFullPath returns the full category path as a string (e.g., "Electronics > Computers > Laptops")
func (c *Category) GetFullPath(db *gorm.DB) string {
	if c.IsParentCategory() {
		return c.Name
	}
	
	var parent Category
	if err := db.First(&parent, c.ParentID).Error; err != nil {
		return c.Name
	}
	
	return parent.GetFullPath(db) + " > " + c.Name
}
