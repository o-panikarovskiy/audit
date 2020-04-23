package user

import "time"

const (
	// AnonymousRole role 1
	AnonymousRole = 1
	// UserRole role 2
	UserRole = 2
	// AdminRole role 4
	AdminRole = 4
)

// User data
type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Created      time.Time `json:"created"`
	Status       string    `json:"-"`
	PasswordHash string    `json:"-"`
	PasswordSalt string    `json:"-"`
	Role         int       `json:"role"`
}
