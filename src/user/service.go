package user

import "audit/src/sessions"

//IService bussines logic interface
type IService interface {
	IReader
	IWriter
	Register(string, string) (*User, string, error)
	Auth(string, string) (*User, string, error)
	CheckSession(string, string) (*User, error)
	GetRepo() IRepository
	GetSessionStorage() sessions.IStorage
	ShutDown()
}
