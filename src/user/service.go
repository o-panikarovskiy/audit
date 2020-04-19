package user

//IService bussines logic interface
type IService interface {
	IReader
	IWriter
	Register(string, string) (*User, error)
	Auth(string, string) (*User, error)
	CheckSession(string) (*User, error)
	GetRepo() IRepository
}
