package user

import "audit/src/sessions"

// IReader interface
type IReader interface {
	FindByID(string) (*User, error)
	FindByUsername(string) (*User, error)
	FindAll() ([]*User, error)
}

// IWriter interface
type IWriter interface {
	Store(*User) (*User, error)
	Update(*User) (*User, error)
}

//IRepository repository interface
type IRepository interface {
	IReader
	IWriter
}

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
