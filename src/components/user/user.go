package user

// User data
type User struct {
	ID           string `json:"id"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"-"`
	PasswordSalt string `json:"-"`
	Role         int    `json:"role"`
}
