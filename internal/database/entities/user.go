package entities

import "github.com/gofrs/uuid"

import "time"

// User represents the data about a given account
type User struct {
	ID            uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:""`
	Username      string
	Email         string
	PasswordHash  string
	ProfileURL    string
	EmailVerified bool
	Role          int
}
