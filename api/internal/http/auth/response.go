package auth

import "time"

type UserResponsePayload struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Image         string    `json:"image"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserWithAccount struct {
	ID            string
	Email         string
	Username      string
	FirstName     string
	LastName      string
	Image         string
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	AccountRowID  string
	UserID        string
	AccountID     string
	ProviderID    string
	PasswordHash  string
}
