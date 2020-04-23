package user

import "audit/src/sessions"

// IService user bussines logic interface
type IService interface {
	IReader
	IWriter
	SignUp(username string, password string) (string, string, error)
	EndSignUp(confirmID string, confirmValue string) (*User, string, error)
	Auth(email string, password string) (*User, string, error)
	RestoreSessionUser(sid string) (*User, error)
	CheckAuthSession(sid string) (*User, error)
	GetRepo() IRepository
	GetSessionStorage() sessions.IStorage
}

// IConfirmService service for confirm user registration
type IConfirmService interface {
	Send(user *User, confirmID string, confirmValue string) error
}
