package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`               // Internal auto-increment ID
	UUID      string    `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"` // Public UUID
	CreatedAt time.Time `json:"created_at"`                                       // Creation timestamp
	UpdatedAt time.Time `json:"updated_at"`                                       // Last update timestamp

	// Associations
	Credentials UserCredentials `gorm:"foreignKey:UserID" json:"credentials"` // One-to-one relationship with credentials
	Profile     UserProfile     `gorm:"foreignKey:UserID" json:"profile"`     // One-to-one relationship with profile
	Addresses   []UserAddress   `gorm:"foreignKey:UserID" json:"addresses"`   // One-to-many relationship with addresses
	Sessions    []UserSession   `gorm:"foreignKey:UserID" json:"sessions"`    // One-to-many relationship with sessions
}

type UserCredentials struct {
	UserID           int    `gorm:"primaryKey" json:"user_id"`               // Foreign key to the users table
	Email            string `gorm:"unique;size:255" json:"email"`            // User email
	PasswordHash     string `gorm:"type:text" json:"password_hash"`          // Hashed password
	TwoFactorEnabled bool   `gorm:"default:false" json:"two_factor_enabled"` // 2FA enabled or not
	TwoFactorSecret  string `gorm:"type:text" json:"two_factor_secret"`      // 2FA secret (nullable)
}

type UserAddress struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`       // Auto-incrementing address ID
	UserID       int    `json:"user_id"`                                  // Foreign key to users table
	AddressLine1 string `gorm:"type:text;not null" json:"address_line_1"` // Address line 1
	AddressLine2 string `gorm:"type:text" json:"address_line_2"`          // Address line 2 (optional)
	City         string `gorm:"size:100;not null" json:"city"`            // City
	State        string `gorm:"size:100;not null" json:"state"`           // State
	Country      string `gorm:"size:100;not null" json:"country"`         // Country
	PostalCode   string `gorm:"size:20;not null" json:"postal_code"`      // Postal/ZIP code
	IsPrimary    bool   `gorm:"default:false" json:"is_primary"`          // Primary address flag
}

type UserProfile struct {
	UserID         int       `gorm:"primaryKey" json:"user_id"`           // Foreign key to users table
	FirstName      string    `gorm:"size:100;not null" json:"first_name"` // First name
	LastName       string    `gorm:"size:100" json:"last_name"`           // Last name (optional)
	DateOfBirth    time.Time `json:"date_of_birth"`                       // Date of birth
	ProfilePicture string    `gorm:"type:text" json:"profile_picture"`    // URL or path to the profile picture
	PhoneNumber    string    `gorm:"size:20" json:"phone_number"`         // Phone number (optional)
}

type UserSession struct {
	SessionID    string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"session_id"` // Unique session ID
	UserID       int       `json:"user_id"`                                                           // Foreign key to users table
	LoginTime    time.Time `gorm:"not null;default:now()" json:"login_time"`                          // Login timestamp
	LastActivity time.Time `gorm:"not null;default:now()" json:"last_activity"`                       // Last activity timestamp
	IPAddress    string    `gorm:"size:45" json:"ip_address"`                                         // IP address
	UserAgent    string    `gorm:"type:text" json:"user_agent"`                                       // Device/browser information
	ExpiresAt    time.Time `gorm:"not null" json:"expires_at"`                                        // Session expiration time
}
