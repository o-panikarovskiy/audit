package emailconfirmator

import (
	"audit/src/config"
	"audit/src/user"
	"log"
)

type emailConfirmatorService struct {
	cfg *config.AppConfig
}

// NewEmailConfirmatorService create new confimator service
func NewEmailConfirmatorService(cfg *config.AppConfig) user.IConfirmService {
	return &emailConfirmatorService{
		cfg: cfg,
	}
}

func (ec *emailConfirmatorService) Send(user *user.User, confirmID string, confirmValue string) error {
	log.Printf("Try to send confirmation email to %s with token %s\n", user.Email, confirmID)
	return nil
}
