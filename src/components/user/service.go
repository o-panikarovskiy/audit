package user

//IService bussines logic interface
type IService interface {
	IReader
	IWriter
	Register(email string, password string) (*User, error)
	Auth(email string, password string) (*User, error)
	CheckSession(sessionID string) (*User, error)
	GetRepo() IRepository
}
