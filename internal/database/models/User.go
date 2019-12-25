package models

// User represents the data about a given account
type User struct {
	BaseModel
	Username      string
	Email         string
	PasswordHash  string
	ProfileURL    string
	EmailVerified bool
	Role          int
}
