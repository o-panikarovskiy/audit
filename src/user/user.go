package user

import "time"

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
