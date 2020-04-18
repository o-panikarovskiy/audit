package user

// User data
type User struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Role  int    `json:"role"`
}
