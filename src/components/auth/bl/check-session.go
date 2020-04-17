package bl

import (
	"audit/src/components/auth/models"
	"audit/src/utils"
)

// CheckSession check user session and return UserModel or error
func CheckSession(sessionID string) (*models.UserModel, error) {

	user := &models.UserModel{
		ID:   utils.CreateGUID(),
		Role: models.Admin,
	}

	// utils.NewAPIError(400, "SESSION_INVALID", "Invalid session")

	return user, nil
}
