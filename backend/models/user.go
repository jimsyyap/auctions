package models

// User represents a user in the system
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"` // Note: This should be hashed before storing in the database
	Role         string `json:"role"`     // Roles: Admin, Seller, Buyer
}
