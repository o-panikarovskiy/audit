package user

//IService bussines logic interface
type IService interface {
	IReader
	IWriter
	Register(user *User) (string, error)
	Auth(user *User, password string) error
	CheckSession(sessionID string) (*User, error)
	GetRepo() IRepository
}
