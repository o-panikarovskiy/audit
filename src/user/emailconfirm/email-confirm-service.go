package emailconfirm

import (
	"audit/src/config"
	"audit/src/user"
	"fmt"
)

type emailConfirmService struct {
	cfg *config.AppConfig
}

// NewEmailConfirmService create new repository
func NewEmailConfirmService(cfg *config.AppConfig) user.IConfirmService {
	return &emailConfirmService{
		cfg: cfg,
	}
}

func (ec *emailConfirmService) Send(user *user.User, confirmID string, confirmValue string) error {
	fmt.Printf("\n\nTry to send confirmation email to %s with token %s\n\n", user.Email, confirmID)
	return nil
}
