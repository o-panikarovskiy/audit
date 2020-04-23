package emailconfirm

import (
	"audit/src/config"
	"audit/src/user"
	"log"
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
	log.Printf("Try to send confirmation email to %s with token %s\n", user.Email, confirmID)
	return nil
}
