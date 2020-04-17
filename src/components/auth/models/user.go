package models

// UserModel model
type UserModel struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Role  int    `json:"role"`
}
