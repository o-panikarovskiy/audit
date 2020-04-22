package user

import "audit/src/sessions"

//IService bussines logic interface
type IService interface {
	IReader
	IWriter
	SignUp(email string, password string) (string, error)
	EndSignUp(tokenID string) (*User, string, error)
	Auth(email string, password string) (*User, string, error)
	RestoreSessionUser(sid string) (*User, error)
	CheckAuthSession(sid string) (*User, error)
	GetRepo() IRepository
	GetSessionStorage() sessions.IStorage
}
